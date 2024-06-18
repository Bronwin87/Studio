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
 * The DataPipelineStatus model module.
 * @module model/DataPipelineStatus
 * @version v1
 */
class DataPipelineStatus {
    /**
     * Constructs a new <code>DataPipelineStatus</code>.
     * @alias module:model/DataPipelineStatus
     */
    constructor() { 
        
        DataPipelineStatus.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>DataPipelineStatus</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DataPipelineStatus} obj Optional instance to populate.
     * @return {module:model/DataPipelineStatus} The populated <code>DataPipelineStatus</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DataPipelineStatus();

            if (data.hasOwnProperty('completed')) {
                obj['completed'] = ApiClient.convertToType(data['completed'], 'Boolean');
            }
            if (data.hasOwnProperty('failed')) {
                obj['failed'] = ApiClient.convertToType(data['failed'], 'Boolean');
            }
            if (data.hasOwnProperty('empty')) {
                obj['empty'] = ApiClient.convertToType(data['empty'], 'Boolean');
            }
            if (data.hasOwnProperty('index')) {
                obj['index'] = ApiClient.convertToType(data['index'], 'String');
            }
            if (data.hasOwnProperty('document_id')) {
                obj['document_id'] = ApiClient.convertToType(data['document_id'], 'String');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], {'String': ['String']});
            }
            if (data.hasOwnProperty('creation')) {
                obj['creation'] = ApiClient.convertToType(data['creation'], 'Date');
            }
            if (data.hasOwnProperty('last_update')) {
                obj['last_update'] = ApiClient.convertToType(data['last_update'], 'Date');
            }
            if (data.hasOwnProperty('steps')) {
                obj['steps'] = ApiClient.convertToType(data['steps'], ['String']);
            }
            if (data.hasOwnProperty('remaining_steps')) {
                obj['remaining_steps'] = ApiClient.convertToType(data['remaining_steps'], ['String']);
            }
            if (data.hasOwnProperty('completed_steps')) {
                obj['completed_steps'] = ApiClient.convertToType(data['completed_steps'], ['String']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>DataPipelineStatus</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>DataPipelineStatus</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['index'] && !(typeof data['index'] === 'string' || data['index'] instanceof String)) {
            throw new Error("Expected the field `index` to be a primitive type in the JSON string but got " + data['index']);
        }
        // ensure the json data is a string
        if (data['document_id'] && !(typeof data['document_id'] === 'string' || data['document_id'] instanceof String)) {
            throw new Error("Expected the field `document_id` to be a primitive type in the JSON string but got " + data['document_id']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['steps'])) {
            throw new Error("Expected the field `steps` to be an array in the JSON data but got " + data['steps']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['remaining_steps'])) {
            throw new Error("Expected the field `remaining_steps` to be an array in the JSON data but got " + data['remaining_steps']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['completed_steps'])) {
            throw new Error("Expected the field `completed_steps` to be an array in the JSON data but got " + data['completed_steps']);
        }

        return true;
    }


}



/**
 * @member {Boolean} completed
 */
DataPipelineStatus.prototype['completed'] = undefined;

/**
 * @member {Boolean} failed
 */
DataPipelineStatus.prototype['failed'] = undefined;

/**
 * @member {Boolean} empty
 */
DataPipelineStatus.prototype['empty'] = undefined;

/**
 * @member {String} index
 */
DataPipelineStatus.prototype['index'] = undefined;

/**
 * @member {String} document_id
 */
DataPipelineStatus.prototype['document_id'] = undefined;

/**
 * @member {Object.<String, Array.<String>>} tags
 */
DataPipelineStatus.prototype['tags'] = undefined;

/**
 * @member {Date} creation
 */
DataPipelineStatus.prototype['creation'] = undefined;

/**
 * @member {Date} last_update
 */
DataPipelineStatus.prototype['last_update'] = undefined;

/**
 * @member {Array.<String>} steps
 */
DataPipelineStatus.prototype['steps'] = undefined;

/**
 * @member {Array.<String>} remaining_steps
 */
DataPipelineStatus.prototype['remaining_steps'] = undefined;

/**
 * @member {Array.<String>} completed_steps
 */
DataPipelineStatus.prototype['completed_steps'] = undefined;






export default DataPipelineStatus;

