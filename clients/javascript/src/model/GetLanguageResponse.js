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
import SupportedLanguage from './SupportedLanguage';

/**
 * The GetLanguageResponse model module.
 * @module model/GetLanguageResponse
 * @version v1
 */
class GetLanguageResponse {
    /**
     * Constructs a new <code>GetLanguageResponse</code>.
     * @alias module:model/GetLanguageResponse
     */
    constructor() { 
        
        GetLanguageResponse.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GetLanguageResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GetLanguageResponse} obj Optional instance to populate.
     * @return {module:model/GetLanguageResponse} The populated <code>GetLanguageResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GetLanguageResponse();

            if (data.hasOwnProperty('languages')) {
                obj['languages'] = ApiClient.convertToType(data['languages'], [SupportedLanguage]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>GetLanguageResponse</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>GetLanguageResponse</code>.
     */
    static validateJSON(data) {
        if (data['languages']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['languages'])) {
                throw new Error("Expected the field `languages` to be an array in the JSON data but got " + data['languages']);
            }
            // validate the optional field `languages` (array)
            for (const item of data['languages']) {
                SupportedLanguage.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {Array.<module:model/SupportedLanguage>} languages
 */
GetLanguageResponse.prototype['languages'] = undefined;






export default GetLanguageResponse;
