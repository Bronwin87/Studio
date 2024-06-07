﻿// Copyright (c) Microsoft. All rights reserved.

using System;
using System.ComponentModel.DataAnnotations;
using System.Reflection;

namespace CopilotChat.WebApi.Options;

/// <summary>
///     If the other property is set to the expected value, then this property is required.
/// </summary>
[AttributeUsage(AttributeTargets.Property)]
internal sealed class RequiredOnPropertyValueAttribute : ValidationAttribute
{
    /// <summary>
    ///     If the other property is set to the expected value, then this property is required.
    /// </summary>
    /// <param name="otherPropertyName">Name of the other property.</param>
    /// <param name="otherPropertyValue">Value of the other property when this property is required.</param>
    /// <param name="notEmptyOrWhitespace">True to make sure that the value is not empty or whitespace when required.</param>
    public RequiredOnPropertyValueAttribute(string otherPropertyName, object? otherPropertyValue,
        bool notEmptyOrWhitespace = true)
    {
        OtherPropertyName = otherPropertyName;
        OtherPropertyValue = otherPropertyValue;
        NotEmptyOrWhitespace = notEmptyOrWhitespace;
    }

    /// <summary>
    ///     Name of the other property.
    /// </summary>
    public string OtherPropertyName { get; }

    /// <summary>
    ///     Value of the other property when this property is required.
    /// </summary>
    public object? OtherPropertyValue { get; }

    /// <summary>
    ///     True to make sure that the value is not empty or whitespace when required.
    /// </summary>
    public bool NotEmptyOrWhitespace { get; }

    protected override ValidationResult? IsValid(object? value, ValidationContext validationContext)
    {
        PropertyInfo? otherPropertyInfo = validationContext.ObjectType.GetRuntimeProperty(OtherPropertyName);

        // If the other property is not found, return an error.
        if (otherPropertyInfo == null)
        {
            return new ValidationResult($"Unknown other property name '{OtherPropertyName}'.");
        }

        // If the other property is an indexer, return an error.
        if (otherPropertyInfo.GetIndexParameters().Length > 0)
        {
            throw new ArgumentException(
                $"Other property not found ('{validationContext.MemberName}, '{OtherPropertyName}').");
        }

        object? otherPropertyValue = otherPropertyInfo.GetValue(validationContext.ObjectInstance, null);

        // If the other property is set to the expected value, then this property is required.
        if (Equals(OtherPropertyValue, otherPropertyValue))
        {
            if (value == null)
            {
                return new ValidationResult(
                    $"Property '{validationContext.DisplayName}' is required when '{OtherPropertyName}' is {OtherPropertyValue}.");
            }

            if (NotEmptyOrWhitespace && string.IsNullOrWhiteSpace(value.ToString()))
            {
                return new ValidationResult(
                    $"Property '{validationContext.DisplayName}' cannot be empty or whitespace when '{OtherPropertyName}' is {OtherPropertyValue}.");
            }

            return ValidationResult.Success;
        }

        return ValidationResult.Success;
    }
}