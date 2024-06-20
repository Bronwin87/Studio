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
import DocumentFilters from './DocumentFilters';

/**
 * The ListDocumentParameters model module.
 * @module model/ListDocumentParameters
 * @version v1
 */
class ListDocumentParameters {
    /**
     * Constructs a new <code>ListDocumentParameters</code>.
     * @alias module:model/ListDocumentParameters
     */
    constructor() { 
        
        ListDocumentParameters.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ListDocumentParameters</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ListDocumentParameters} obj Optional instance to populate.
     * @return {module:model/ListDocumentParameters} The populated <code>ListDocumentParameters</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ListDocumentParameters();

            if (data.hasOwnProperty('index')) {
                obj['index'] = ApiClient.convertToType(data['index'], 'String');
            }
            if (data.hasOwnProperty('filter')) {
                obj['filter'] = ApiClient.convertToType(data['filter'], [DocumentFilters]);
            }
            if (data.hasOwnProperty('withEmbeddings')) {
                obj['withEmbeddings'] = ApiClient.convertToType(data['withEmbeddings'], 'Boolean');
            }
            if (data.hasOwnProperty('limit')) {
                obj['limit'] = ApiClient.convertToType(data['limit'], 'Number');
            }
            if (data.hasOwnProperty('offset')) {
                obj['offset'] = ApiClient.convertToType(data['offset'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ListDocumentParameters</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ListDocumentParameters</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['index'] && !(typeof data['index'] === 'string' || data['index'] instanceof String)) {
            throw new Error("Expected the field `index` to be a primitive type in the JSON string but got " + data['index']);
        }
        if (data['filter']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['filter'])) {
                throw new Error("Expected the field `filter` to be an array in the JSON data but got " + data['filter']);
            }
            // validate the optional field `filter` (array)
            for (const item of data['filter']) {
                DocumentFilters.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * Optional index to specify which index to search in. Defaults to 'default'
 * @member {String} index
 */
ListDocumentParameters.prototype['index'] = undefined;

/**
 * Optional filtering of document id(s) and/or tags
 * @member {Array.<module:model/DocumentFilters>} filter
 */
ListDocumentParameters.prototype['filter'] = undefined;

/**
 * Optionally specifies if embedding should be returned in response. Default is false
 * @member {Boolean} withEmbeddings
 */
ListDocumentParameters.prototype['withEmbeddings'] = undefined;

/**
 * Optional filter for specifying maximum number of entries to return. Defaults to 3
 * @member {Number} limit
 */
ListDocumentParameters.prototype['limit'] = undefined;

/**
 * Optional filter for specifying list offset for paging. Default is 0
 * @member {Number} offset
 */
ListDocumentParameters.prototype['offset'] = undefined;






export default ListDocumentParameters;
