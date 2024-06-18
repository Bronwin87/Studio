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

import ApiClient from '../ApiClient';

/**
 * The ReRankDocumentsRequest model module.
 * @module model/ReRankDocumentsRequest
 * @version v1
 */
class ReRankDocumentsRequest {
    /**
     * Constructs a new <code>ReRankDocumentsRequest</code>.
     * @alias module:model/ReRankDocumentsRequest
     */
    constructor() { 
        
        ReRankDocumentsRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ReRankDocumentsRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ReRankDocumentsRequest} obj Optional instance to populate.
     * @return {module:model/ReRankDocumentsRequest} The populated <code>ReRankDocumentsRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ReRankDocumentsRequest();

            if (data.hasOwnProperty('model')) {
                obj['model'] = ApiClient.convertToType(data['model'], 'String');
            }
            if (data.hasOwnProperty('query')) {
                obj['query'] = ApiClient.convertToType(data['query'], 'String');
            }
            if (data.hasOwnProperty('documents')) {
                obj['documents'] = ApiClient.convertToType(data['documents'], ['String']);
            }
            if (data.hasOwnProperty('limit')) {
                obj['limit'] = ApiClient.convertToType(data['limit'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ReRankDocumentsRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ReRankDocumentsRequest</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['model'] && !(typeof data['model'] === 'string' || data['model'] instanceof String)) {
            throw new Error("Expected the field `model` to be a primitive type in the JSON string but got " + data['model']);
        }
        // ensure the json data is a string
        if (data['query'] && !(typeof data['query'] === 'string' || data['query'] instanceof String)) {
            throw new Error("Expected the field `query` to be a primitive type in the JSON string but got " + data['query']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['documents'])) {
            throw new Error("Expected the field `documents` to be an array in the JSON data but got " + data['documents']);
        }

        return true;
    }


}



/**
 * Optional embedding model. Defaults to configured default
 * @member {String} model
 */
ReRankDocumentsRequest.prototype['model'] = undefined;

/**
 * Semantic query to find matching documents
 * @member {String} query
 */
ReRankDocumentsRequest.prototype['query'] = undefined;

/**
 * Optional index to specify which index to search in. Defaults to 'default'
 * @member {Array.<String>} documents
 */
ReRankDocumentsRequest.prototype['documents'] = undefined;

/**
 * Optional filter for specifying maximum number of entries to return. Defaults to 3
 * @member {Number} limit
 */
ReRankDocumentsRequest.prototype['limit'] = undefined;






export default ReRankDocumentsRequest;

