/**
 * Studio - AI Empower Labs
 * # Studio API Documentation  ## Introduction Welcome to Studio by AI Empower Labs API documentation! We are thrilled to offer developers around the world access to our cutting-edge artificial intelligence technology and semantic search. Our API is designed to empower your applications with state-of-the-art AI capabilities, including but not limited to natural language processing, audio transcription, embedding, and predictive analytics.  Our mission is to make AI technology accessible and easy to integrate, enabling you to enhance your applications, improve user experiences, and innovate in your field. Whether you're building complex systems, developing mobile apps, or creating web services, our API provides you with the tools you need to incorporate AI functionalities seamlessly.  Support and Feedback We are committed to providing exceptional support to our developer community. If you have any questions, encounter any issues, or have feedback on how we can improve our API, please don't hesitate to contact our support team @ support@AIEmpowerLabs.com.  Terms of Use Please review our terms of use and privacy policy before integrating our API into your application. By using our API, you agree to comply with these terms.
 *
 * The version of the OpenAPI document: v1
 * Contact: support@aiempowerlabs.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import DataPipelineStatus from '../model/DataPipelineStatus';
import DeleteAccepted from '../model/DeleteAccepted';
import IndexCollection from '../model/IndexCollection';
import MemoryAnswer from '../model/MemoryAnswer';
import MemoryQuery from '../model/MemoryQuery';
import ProblemDetails from '../model/ProblemDetails';
import SearchQuery from '../model/SearchQuery';
import SearchResult from '../model/SearchResult';
import StreamableFileContent from '../model/StreamableFileContent';
import UploadAccepted from '../model/UploadAccepted';

/**
* KernelMemory service.
* @module api/KernelMemoryApi
* @version v1
*/
export default class KernelMemoryApi {

    /**
    * Constructs a new KernelMemoryApi. 
    * @alias module:api/KernelMemoryApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the downloadGet operation.
     * @callback module:api/KernelMemoryApi~downloadGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/StreamableFileContent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} documentId 
     * @param {String} filename 
     * @param {Object} opts Optional parameters
     * @param {String} [index] 
     * @param {module:api/KernelMemoryApi~downloadGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/StreamableFileContent}
     */
    downloadGet(documentId, filename, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'documentId' is set
      if (documentId === undefined || documentId === null) {
        throw new Error("Missing the required parameter 'documentId' when calling downloadGet");
      }
      // verify the required parameter 'filename' is set
      if (filename === undefined || filename === null) {
        throw new Error("Missing the required parameter 'filename' when calling downloadGet");
      }

      let pathParams = {
      };
      let queryParams = {
        'index': opts['index'],
        'documentId': documentId,
        'filename': filename
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json', 'application/problem+json'];
      let returnType = StreamableFileContent;
      return this.apiClient.callApi(
        '/download', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the kernelMemoryAsk operation.
     * @callback module:api/KernelMemoryApi~kernelMemoryAskCallback
     * @param {String} error Error message, if any.
     * @param {module:model/MemoryAnswer} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Query documents and forward result to LLM
     * @param {module:model/MemoryQuery} memoryQuery 
     * @param {module:api/KernelMemoryApi~kernelMemoryAskCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/MemoryAnswer}
     */
    kernelMemoryAsk(memoryQuery, callback) {
      let postBody = memoryQuery;
      // verify the required parameter 'memoryQuery' is set
      if (memoryQuery === undefined || memoryQuery === null) {
        throw new Error("Missing the required parameter 'memoryQuery' when calling kernelMemoryAsk");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json', 'application/problem+json'];
      let returnType = MemoryAnswer;
      return this.apiClient.callApi(
        '/api/kernelmemory/ask', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the kernelMemoryDelete operation.
     * @callback module:api/KernelMemoryApi~kernelMemoryDeleteCallback
     * @param {String} error Error message, if any.
     * @param {module:model/DeleteAccepted} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete document from specific index
     * @param {String} documentId 
     * @param {Object} opts Optional parameters
     * @param {String} [index] 
     * @param {module:api/KernelMemoryApi~kernelMemoryDeleteCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/DeleteAccepted}
     */
    kernelMemoryDelete(documentId, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'documentId' is set
      if (documentId === undefined || documentId === null) {
        throw new Error("Missing the required parameter 'documentId' when calling kernelMemoryDelete");
      }

      let pathParams = {
      };
      let queryParams = {
        'index': opts['index'],
        'documentId': documentId
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json', 'application/problem+json'];
      let returnType = DeleteAccepted;
      return this.apiClient.callApi(
        '/api/kernelmemory/documents', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the kernelMemoryDeleteIndex operation.
     * @callback module:api/KernelMemoryApi~kernelMemoryDeleteIndexCallback
     * @param {String} error Error message, if any.
     * @param {module:model/DeleteAccepted} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete index
     * @param {Object} opts Optional parameters
     * @param {String} [index] 
     * @param {module:api/KernelMemoryApi~kernelMemoryDeleteIndexCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/DeleteAccepted}
     */
    kernelMemoryDeleteIndex(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'index': opts['index']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json', 'application/problem+json'];
      let returnType = DeleteAccepted;
      return this.apiClient.callApi(
        '/api/kernelmemory/indexes', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the kernelMemoryGetIndex operation.
     * @callback module:api/KernelMemoryApi~kernelMemoryGetIndexCallback
     * @param {String} error Error message, if any.
     * @param {module:model/IndexCollection} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List indexes
     * @param {module:api/KernelMemoryApi~kernelMemoryGetIndexCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/IndexCollection}
     */
    kernelMemoryGetIndex(callback) {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json', 'application/problem+json'];
      let returnType = IndexCollection;
      return this.apiClient.callApi(
        '/api/kernelmemory/indexes', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the kernelMemorySearch operation.
     * @callback module:api/KernelMemoryApi~kernelMemorySearchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/SearchResult} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Search for documents in specific index
     * @param {module:model/SearchQuery} searchQuery 
     * @param {module:api/KernelMemoryApi~kernelMemorySearchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/SearchResult}
     */
    kernelMemorySearch(searchQuery, callback) {
      let postBody = searchQuery;
      // verify the required parameter 'searchQuery' is set
      if (searchQuery === undefined || searchQuery === null) {
        throw new Error("Missing the required parameter 'searchQuery' when calling kernelMemorySearch");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json', 'application/problem+json'];
      let returnType = SearchResult;
      return this.apiClient.callApi(
        '/api/kernelmemory/search', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the kernelMemoryUpload operation.
     * @callback module:api/KernelMemoryApi~kernelMemoryUploadCallback
     * @param {String} error Error message, if any.
     * @param {module:model/UploadAccepted} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Upload file for ingestion
     * @param {module:api/KernelMemoryApi~kernelMemoryUploadCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/UploadAccepted}
     */
    kernelMemoryUpload(callback) {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json', 'application/problem+json'];
      let returnType = UploadAccepted;
      return this.apiClient.callApi(
        '/api/kernelmemory/upload', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the kernelMemoryUploadStatus operation.
     * @callback module:api/KernelMemoryApi~kernelMemoryUploadStatusCallback
     * @param {String} error Error message, if any.
     * @param {module:model/DataPipelineStatus} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get ingestion status for specific document
     * @param {String} documentId 
     * @param {Object} opts Optional parameters
     * @param {String} [index] 
     * @param {module:api/KernelMemoryApi~kernelMemoryUploadStatusCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/DataPipelineStatus}
     */
    kernelMemoryUploadStatus(documentId, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'documentId' is set
      if (documentId === undefined || documentId === null) {
        throw new Error("Missing the required parameter 'documentId' when calling kernelMemoryUploadStatus");
      }

      let pathParams = {
      };
      let queryParams = {
        'index': opts['index'],
        'documentId': documentId
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json', 'application/problem+json'];
      let returnType = DataPipelineStatus;
      return this.apiClient.callApi(
        '/api/kernelmemory/upload-status', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
