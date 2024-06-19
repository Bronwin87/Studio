/*
 * Studio - AI Empower Labs
 * # Studio API Documentation  ## Introduction Welcome to Studio by AI Empower Labs API documentation! We are thrilled to offer developers around the world access to our cutting-edge artificial intelligence technology and semantic search. Our API is designed to empower your applications with state-of-the-art AI capabilities, including but not limited to natural language processing, audio transcription, embedding, and predictive analytics.  Our mission is to make AI technology accessible and easy to integrate, enabling you to enhance your applications, improve user experiences, and innovate in your field. Whether you're building complex systems, developing mobile apps, or creating web services, our API provides you with the tools you need to incorporate AI functionalities seamlessly.  Support and Feedback We are committed to providing exceptional support to our developer community. If you have any questions, encounter any issues, or have feedback on how we can improve our API, please don't hesitate to contact our support team @ support@AIEmpowerLabs.com.  Terms of Use Please review our terms of use and privacy policy before integrating our API into your application. By using our API, you agree to comply with these terms.
 *
 * The version of the OpenAPI document: v1
 * Contact: support@aiempowerlabs.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.AskDocumentRequest;
import org.openapitools.client.model.AskDocumentResponse;
import org.openapitools.client.model.DataPipelineStatus;
import java.io.File;
import org.openapitools.client.model.HttpValidationProblemDetails;
import org.openapitools.client.model.IngestDocumentResponse;
import org.openapitools.client.model.IngestTextDocumentRequest;
import org.openapitools.client.model.IngestWebPageDocumentRequest;
import org.openapitools.client.model.ListDocumentParameters;
import org.openapitools.client.model.ListDocumentResponse;
import org.openapitools.client.model.ProblemDetails;
import org.openapitools.client.model.QueryDocumentRequest;
import org.openapitools.client.model.QueryDocumentResponse;
import org.openapitools.client.model.ReRankDocumentsRequest;
import org.openapitools.client.model.ReRankDocumentsResponse;
import java.util.UUID;
import org.openapitools.client.model.WordCloudDocumentRequest;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for SemanticSearchApi
 */
@Disabled
public class SemanticSearchApiTest {

    private final SemanticSearchApi api = new SemanticSearchApi();

    /**
     * Ask questions over ingested documents
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchAskTest() throws ApiException {
        AskDocumentRequest askDocumentRequest = null;
        AskDocumentResponse response = api.semanticSearchAsk(askDocumentRequest);
        // TODO: test validations
    }

    /**
     * Delete specific document by id
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchDeleteDocumentTest() throws ApiException {
        String documentId = null;
        String index = null;
        api.semanticSearchDeleteDocument(documentId, index);
        // TODO: test validations
    }

    /**
     * Delete specific index by name
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchDeleteIndexTest() throws ApiException {
        String name = null;
        api.semanticSearchDeleteIndex(name);
        // TODO: test validations
    }

    /**
     * Import file document into semantic search
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchFileIngestionTest() throws ApiException {
        List<File> files = null;
        String documentId = null;
        String index = null;
        List<String> pipeline = null;
        String webHookUrl = null;
        String embeddingModel = null;
        Map<String, Object> args = null;
        Map<String, Object> tags = null;
        IngestDocumentResponse response = api.semanticSearchFileIngestion(files, documentId, index, pipeline, webHookUrl, embeddingModel, args, tags);
        // TODO: test validations
    }

    /**
     * Get queue status for ingestion job
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchIngestionStatusTest() throws ApiException {
        UUID id = null;
        DataPipelineStatus response = api.semanticSearchIngestionStatus(id);
        // TODO: test validations
    }

    /**
     * List - and filter - for ingested documents
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchListTest() throws ApiException {
        ListDocumentParameters listDocumentParameters = null;
        ListDocumentResponse response = api.semanticSearchList(listDocumentParameters);
        // TODO: test validations
    }

    /**
     * Query ingested documents using semantic search
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchQueryTest() throws ApiException {
        QueryDocumentRequest queryDocumentRequest = null;
        QueryDocumentResponse response = api.semanticSearchQuery(queryDocumentRequest);
        // TODO: test validations
    }

    /**
     * Rerank documents
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchRerankTest() throws ApiException {
        ReRankDocumentsRequest reRankDocumentsRequest = null;
        ReRankDocumentsResponse response = api.semanticSearchRerank(reRankDocumentsRequest);
        // TODO: test validations
    }

    /**
     * Import plain text into semantic search
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchTextIngestionTest() throws ApiException {
        IngestTextDocumentRequest ingestTextDocumentRequest = null;
        IngestDocumentResponse response = api.semanticSearchTextIngestion(ingestTextDocumentRequest);
        // TODO: test validations
    }

    /**
     * Import web page text into semantic search
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchWebpageIngestionTest() throws ApiException {
        IngestWebPageDocumentRequest ingestWebPageDocumentRequest = null;
        IngestDocumentResponse response = api.semanticSearchWebpageIngestion(ingestWebPageDocumentRequest);
        // TODO: test validations
    }

    /**
     * Generate word cloud from semantic database
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void semanticSearchWordcloudTest() throws ApiException {
        WordCloudDocumentRequest wordCloudDocumentRequest = null;
        byte[] response = api.semanticSearchWordcloud(wordCloudDocumentRequest);
        // TODO: test validations
    }

}
