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
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Org.OpenAPITools.Client.OpenAPIDateConverter;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// NamedEntityRecognitionRequest
    /// </summary>
    [DataContract(Name = "NamedEntityRecognitionRequest")]
    public partial class NamedEntityRecognitionRequest : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="NamedEntityRecognitionRequest" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected NamedEntityRecognitionRequest() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="NamedEntityRecognitionRequest" /> class.
        /// </summary>
        /// <param name="text">Text to perform Named Entity Recognition on  It recognises Persons, Organisations, Dates, Locations and Emails (required).</param>
        /// <param name="model">LLM Model to use.</param>
        /// <param name="entities">Entities to extract (required).</param>
        public NamedEntityRecognitionRequest(string text = default(string), string model = default(string), List<string> entities = default(List<string>))
        {
            // to ensure "text" is required (not null)
            if (text == null)
            {
                throw new ArgumentNullException("text is a required property for NamedEntityRecognitionRequest and cannot be null");
            }
            this.Text = text;
            // to ensure "entities" is required (not null)
            if (entities == null)
            {
                throw new ArgumentNullException("entities is a required property for NamedEntityRecognitionRequest and cannot be null");
            }
            this.Entities = entities;
            this.Model = model;
        }

        /// <summary>
        /// Text to perform Named Entity Recognition on  It recognises Persons, Organisations, Dates, Locations and Emails
        /// </summary>
        /// <value>Text to perform Named Entity Recognition on  It recognises Persons, Organisations, Dates, Locations and Emails</value>
        [DataMember(Name = "text", IsRequired = true, EmitDefaultValue = true)]
        public string Text { get; set; }

        /// <summary>
        /// LLM Model to use
        /// </summary>
        /// <value>LLM Model to use</value>
        [DataMember(Name = "model", EmitDefaultValue = true)]
        public string Model { get; set; }

        /// <summary>
        /// Entities to extract
        /// </summary>
        /// <value>Entities to extract</value>
        [DataMember(Name = "entities", IsRequired = true, EmitDefaultValue = true)]
        public List<string> Entities { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class NamedEntityRecognitionRequest {\n");
            sb.Append("  Text: ").Append(Text).Append("\n");
            sb.Append("  Model: ").Append(Model).Append("\n");
            sb.Append("  Entities: ").Append(Entities).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            // Model (string) maxLength
            if (this.Model != null && this.Model.Length > 250)
            {
                yield return new ValidationResult("Invalid value for Model, length must be less than 250.", new [] { "Model" });
            }

            // Model (string) minLength
            if (this.Model != null && this.Model.Length < 1)
            {
                yield return new ValidationResult("Invalid value for Model, length must be greater than 1.", new [] { "Model" });
            }

            yield break;
        }
    }

}
