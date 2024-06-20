/*
 * Studio - AI Empower Labs
 *
 * # Studio API Documentation  ## Introduction Welcome to Studio by AI Empower Labs API documentation! We are thrilled to offer developers around the world access to our cutting-edge artificial intelligence technology and semantic search. Our API is designed to empower your applications with state-of-the-art AI capabilities, including but not limited to natural language processing, audio transcription, embedding, and predictive analytics.  Our mission is to make AI technology accessible and easy to integrate, enabling you to enhance your applications, improve user experiences, and innovate in your field. Whether you're building complex systems, developing mobile apps, or creating web services, our API provides you with the tools you need to incorporate AI functionalities seamlessly.  Support and Feedback We are committed to providing exceptional support to our developer community. If you have any questions, encounter any issues, or have feedback on how we can improve our API, please don't hesitate to contact our support team @ support@AIEmpowerLabs.com.  Terms of Use Please review our terms of use and privacy policy before integrating our API into your application. By using our API, you agree to comply with these terms.
 *
 * The version of the OpenAPI document: v1
 * Contact: support@aiempowerlabs.com
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */

using System;
using System.IO;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Reflection;
using RestSharp;
using Xunit;

using Org.OpenAPITools.Client;
using Org.OpenAPITools.Api;
// uncomment below to import models
//using Org.OpenAPITools.Model;

namespace Org.OpenAPITools.Test.Api
{
    /// <summary>
    ///  Class for testing SemanticSearchApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    public class SemanticSearchApiTests : IDisposable
    {
        private SemanticSearchApi instance;

        public SemanticSearchApiTests()
        {
            instance = new SemanticSearchApi();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of SemanticSearchApi
        /// </summary>
        [Fact]
        public void InstanceTest()
        {
            // TODO uncomment below to test 'IsType' SemanticSearchApi
            //Assert.IsType<SemanticSearchApi>(instance);
        }

        /// <summary>
        /// Test SemanticSearchAsk
        /// </summary>
        [Fact]
        public void SemanticSearchAskTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //AskDocumentRequest askDocumentRequest = null;
            //var response = instance.SemanticSearchAsk(askDocumentRequest);
            //Assert.IsType<AskDocumentResponse>(response);
        }

        /// <summary>
        /// Test SemanticSearchDeleteDocument
        /// </summary>
        [Fact]
        public void SemanticSearchDeleteDocumentTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string documentId = null;
            //string index = null;
            //instance.SemanticSearchDeleteDocument(documentId, index);
        }

        /// <summary>
        /// Test SemanticSearchDeleteIndex
        /// </summary>
        [Fact]
        public void SemanticSearchDeleteIndexTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string name = null;
            //instance.SemanticSearchDeleteIndex(name);
        }

        /// <summary>
        /// Test SemanticSearchFileIngestion
        /// </summary>
        [Fact]
        public void SemanticSearchFileIngestionTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //List<System.IO.Stream> files = null;
            //string? documentId = null;
            //string? index = null;
            //List<string>? pipeline = null;
            //string? webHookUrl = null;
            //string? embeddingModel = null;
            //Dictionary<string, Object>? args = null;
            //Dictionary<string, Object>? tags = null;
            //var response = instance.SemanticSearchFileIngestion(files, documentId, index, pipeline, webHookUrl, embeddingModel, args, tags);
            //Assert.IsType<IngestDocumentResponse>(response);
        }

        /// <summary>
        /// Test SemanticSearchIngestionStatus
        /// </summary>
        [Fact]
        public void SemanticSearchIngestionStatusTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //Guid id = null;
            //var response = instance.SemanticSearchIngestionStatus(id);
            //Assert.IsType<DataPipelineStatus>(response);
        }

        /// <summary>
        /// Test SemanticSearchList
        /// </summary>
        [Fact]
        public void SemanticSearchListTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //ListDocumentParameters listDocumentParameters = null;
            //var response = instance.SemanticSearchList(listDocumentParameters);
            //Assert.IsType<ListDocumentResponse>(response);
        }

        /// <summary>
        /// Test SemanticSearchQuery
        /// </summary>
        [Fact]
        public void SemanticSearchQueryTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //QueryDocumentRequest queryDocumentRequest = null;
            //var response = instance.SemanticSearchQuery(queryDocumentRequest);
            //Assert.IsType<QueryDocumentResponse>(response);
        }

        /// <summary>
        /// Test SemanticSearchRerank
        /// </summary>
        [Fact]
        public void SemanticSearchRerankTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //ReRankDocumentsRequest reRankDocumentsRequest = null;
            //var response = instance.SemanticSearchRerank(reRankDocumentsRequest);
            //Assert.IsType<ReRankDocumentsResponse>(response);
        }

        /// <summary>
        /// Test SemanticSearchTextIngestion
        /// </summary>
        [Fact]
        public void SemanticSearchTextIngestionTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //IngestTextDocumentRequest ingestTextDocumentRequest = null;
            //var response = instance.SemanticSearchTextIngestion(ingestTextDocumentRequest);
            //Assert.IsType<IngestDocumentResponse>(response);
        }

        /// <summary>
        /// Test SemanticSearchWebpageIngestion
        /// </summary>
        [Fact]
        public void SemanticSearchWebpageIngestionTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //IngestWebPageDocumentRequest ingestWebPageDocumentRequest = null;
            //var response = instance.SemanticSearchWebpageIngestion(ingestWebPageDocumentRequest);
            //Assert.IsType<IngestDocumentResponse>(response);
        }
    }
}