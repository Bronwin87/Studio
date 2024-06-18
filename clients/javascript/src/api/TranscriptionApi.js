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
import HttpValidationProblemDetails from '../model/HttpValidationProblemDetails';
import ProblemDetails from '../model/ProblemDetails';
import TranscriptionAudioUploadResult from '../model/TranscriptionAudioUploadResult';

/**
* Transcription service.
* @module api/TranscriptionApi
* @version v1
*/
export default class TranscriptionApi {

    /**
    * Constructs a new TranscriptionApi. 
    * @alias module:api/TranscriptionApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the transcriptionAsynchronous operation.
     * @callback module:api/TranscriptionApi~transcriptionAsynchronousCallback
     * @param {String} error Error message, if any.
     * @param {module:model/TranscriptionAudioUploadResult} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Upload audio file for asynchronous transcription
     * This service is designed to convert spoken words from audio or video files into written text, utilizing sophisticated speech recognition algorithms for accurate transcription. It offers a range of features that cater to various needs and use cases, making it particularly valuable for journalists, researchers, podcasters, and professionals dealing with meetings, interviews, lectures, or presentations.  ### Key features and capabilities include:  * Support for Various File Formats: Accepts a wide range of audio and video file formats, ensuring flexibility in file uploads. * Advanced Processing Steps: Incorporates noise reduction, speaker diarization, and speech-to-text conversion for clear and differentiated transcriptions. * Asynchronous Background Processing: Allows for non-blocking, efficient handling of transcription tasks, suitable for large files or batches of files. * Webhook Callback URL: Offers real-time updates on the transcription process via a provided webhook, enabling immediate reaction to task completion or failure. * /api/transcribe GET Endpoint: Provides an alternative for users to manually check the status of their transcription requests, allowing flexibility in monitoring. * Automatic Text Translation Feature: An optional service that translates the transcribed text into a specified target language, enhancing the utility for multi-lingual contexts. * Multi-File and Multi-Channel Support: Supports concurrent file uploads and accurate transcription of multi-channel recordings, ideal for complex audio environments. * The transcription output is meticulously formatted to clearly distinguish between channels and speakers, including timestamps and labels for easy navigation and reference. This structured approach ensures that even in challenging audio environments with multiple speakers or channels, the transcription service can provide clear, accurate, and useful text representations of the spoken content.  This service integrates into applications via API calls, offering developers a powerful tool to enhance their applications with audio-to-text conversion capabilities. The inclusion of features like language detection, support for multiple languages, and customization options for specific vocabulary or industry terms further extends its applicability across various domains and industries.
     * @param {Array.<File>} files The file object to ingest.
     * @param {Object} opts Optional parameters
     * @param {module:model/String} [model] Model to use for transcription (Optional, default = Base)
     * @param {String} [language] The language of the input audio. Supplying the input language in [ISO-639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) format will improve accuracy and latency.  (optional)
     * @param {String} [prompt] An optional text to guide the model's style or continue a previous audio segment. The prompt should match the audio language.  (optional)
     * @param {Number} [temperature = 0)] The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use [log probability](https://en.wikipedia.org/wiki/Log_probability) to automatically increase the temperature until certain thresholds are hit.  (optional, default to 0M)
     * @param {String} [webHookUrl] Url to call when transcription has completed or failed. (optional)
     * @param {String} [translateTo] The language to translate transcription into. Supplying the language in [ISO-639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) format will improve accuracy and latency.  (optional)
     * @param {Boolean} [splitOnWord = false)] Split into word segments. (optional, default is false)
     * @param {Boolean} [languageDetection = false)] Enable transcription language detection (Optional. default is false)
     * @param {Boolean} [noiseReduction = false)] Enable noise reduction from audio stream before transcription (Optional. default is false)
     * @param {module:api/TranscriptionApi~transcriptionAsynchronousCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/TranscriptionAudioUploadResult}
     */
    transcriptionAsynchronous(files, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'files' is set
      if (files === undefined || files === null) {
        throw new Error("Missing the required parameter 'files' when calling transcriptionAsynchronous");
      }

      let pathParams = {
      };
      let queryParams = {
        'model': opts['model'],
        'language': opts['language'],
        'prompt': opts['prompt'],
        'temperature': opts['temperature'],
        'webHookUrl': opts['webHookUrl'],
        'translateTo': opts['translateTo'],
        'splitOnWord': opts['splitOnWord'],
        'languageDetection': opts['languageDetection'],
        'noiseReduction': opts['noiseReduction']
      };
      let headerParams = {
      };
      let formParams = {
        'files': this.apiClient.buildCollectionParam(files, 'passthrough')
      };

      let authNames = [];
      let contentTypes = ['multipart/form-data', 'application/x-www-form-urlencoded'];
      let accepts = ['application/json', 'application/problem+json'];
      let returnType = TranscriptionAudioUploadResult;
      return this.apiClient.callApi(
        '/api/transcribe/upload', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the transcriptionGetById operation.
     * @callback module:api/TranscriptionApi~transcriptionGetByIdCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get transcription status and data
     * The /api/transcribe GET endpoint is a crucial component of the audio transcription service, designed to offer users a way to check the status of their transcription requests. This endpoint caters to the needs of users who prefer polling to monitor their requests over relying on webhook callbacks for real-time updates. Here's a detailed description of its functionality and how it integrates within the service:  ### Purpose and Functionality The primary purpose of the /api/transcribe GET endpoint is to provide users with the ability to manually inquire about the current status of their audio or video file transcription tasks. This endpoint supports a polling mechanism, where users can send a GET request at their convenience to receive the latest update on their transcription process.  ### How It Works Request: To utilize this endpoint, users initiate a GET request, including a unique identifier for the transcription task as a parameter. This identifier is provided by the service when the transcription request is first submitted. Response: In response to the GET request, the endpoint returns data about the transcription task's status. The response indicate that the transcription is still processing, has been completed, or has failed.  ### Response Details The response from the /api/transcribe GET endpoint includes several pieces of information that are crucial for users to understand the status and outcome of their transcription requests:  Status: Indicates the current state of the transcription task (e.g., Queued, Completed, Failed). Completion Details: If the transcription is completed, the response include details the resulting transcript.  ### Use Cases This endpoint is particularly useful for scenarios where users need or prefer to periodically check the status of their requests rather than implement real-time update mechanisms via webhooks. It provides flexibility in handling transcription tasks, allowing users to:  ### Advantages The /api/transcribe GET endpoint offers several advantages, including simplicity in implementation, flexibility in usage, and the ability to integrate easily into various application workflows. It provides a straightforward method for users to remain informed about their transcription tasks without the need for complex callback systems, making it an essential feature for many applications and services requiring transcription capabilities.
     * @param {String} id 
     * @param {module:api/TranscriptionApi~transcriptionGetByIdCallback} callback The callback function, accepting three arguments: error, data, response
     */
    transcriptionGetById(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling transcriptionGetById");
      }

      let pathParams = {
      };
      let queryParams = {
        'id': id
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/problem+json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/transcribe', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
