﻿// Copyright (c) Microsoft. All rights reserved.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using CopilotChat.WebApi.Auth;
using CopilotChat.WebApi.Extensions;
using CopilotChat.WebApi.Models.Request;
using CopilotChat.WebApi.Options;
using CopilotChat.WebApi.Storage;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using Microsoft.KernelMemory;

namespace CopilotChat.WebApi.Controllers;

/// <summary>
///     Controller for retrieving kernel memory data of chat sessions.
/// </summary>
[ApiController]
public sealed class ChatMemoryController(
    ILogger<ChatMemoryController> logger,
    IOptions<PromptsOptions> promptsOptions,
    ChatSessionRepository chatSessionRepository) : ControllerBase
{
    /// <summary>
    ///     Gets the kernel memory for the chat session.
    /// </summary>
    /// <param name="memoryClient">The semantic text memory instance.</param>
    /// <param name="chatId">The chat id.</param>
    /// <param name="type">Type of memory. Must map to a member of <see cref="SemanticMemoryType" />.</param>
    [HttpGet]
    [Route("chats/{chatId:guid}/memories")]
    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [Authorize(Policy = AuthPolicyName.RequireChatParticipant)]
    public async Task<IActionResult> GetSemanticMemoriesAsync(
        [FromServices] IKernelMemory memoryClient,
        [FromRoute] string chatId,
        [FromQuery] string type)
    {
        // Sanitize the log input by removing new line characters.
        // https://github.com/microsoft/chat-copilot/security/code-scanning/1
        string sanitizedChatId = GetSanitizedParameter(chatId);
        string sanitizedMemoryType = GetSanitizedParameter(type);

        // Map the requested memoryType to the memory store container name
        if (!promptsOptions.Value.TryGetMemoryContainerName(type, out string memoryContainerName))
        {
            logger.LogWarning("Memory type: {MemoryType} is invalid", sanitizedMemoryType);
            return BadRequest($"Memory type: {sanitizedMemoryType} is invalid.");
        }

        // Make sure the chat session exists.
        if (!await chatSessionRepository.TryFindById(chatId))
        {
            logger.LogWarning("Chat session: {ChatId} does not exist", sanitizedChatId);
            return BadRequest($"Chat session: {sanitizedChatId} does not exist.");
        }

        // Gather the requested kernel memory.
        // Will use a dummy query since we don't care about relevance.
        // minRelevanceScore is set to 0.0 to return all memories.
        List<string> memories = new();
        try
        {
            // Search if there is already a memory item that has a high similarity score with the new item.
            MemoryFilter filter = new();
            filter.ByTag("chatid", chatId);
            filter.ByTag("memory", memoryContainerName);

            SearchResult searchResult =
                await memoryClient.SearchMemoryAsync(
                    promptsOptions.Value.MemoryIndexName,
                    "*",
                    0,
                    1,
                    chatId,
                    memoryContainerName);

            foreach (Citation.Partition memory in searchResult.Results.SelectMany(c => c.Partitions))
            {
                memories.Add(memory.Text);
            }
        }
        catch (Exception connectorException) when (!connectorException.IsCriticalException())
        {
            // A store exception might be thrown if the collection does not exist, depending on the memory store connector.
            logger.LogError(connectorException, "Cannot search collection {ContainerName}", memoryContainerName);
        }

        return Ok(memories);
    }

    #region Private

    private static string GetSanitizedParameter(string parameterValue)
    {
        return parameterValue.Replace(Environment.NewLine, string.Empty, StringComparison.Ordinal);
    }

    # endregion
}
