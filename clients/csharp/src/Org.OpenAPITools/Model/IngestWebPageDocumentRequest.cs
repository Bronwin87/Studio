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
    /// IngestWebPageDocumentRequest
    /// </summary>
    [DataContract(Name = "IngestWebPageDocumentRequest")]
    public partial class IngestWebPageDocumentRequest : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="IngestWebPageDocumentRequest" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected IngestWebPageDocumentRequest() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="IngestWebPageDocumentRequest" /> class.
        /// </summary>
        /// <param name="documentId">Id that uniquely identifies content. Previously ingested documents with the same id will be overwritten (required).</param>
        /// <param name="index">Optional value to specify with index the document should be ingested. Defaults to &#39;default&#39;.</param>
        /// <param name="tags">Optionally add tags to ingestion.</param>
        /// <param name="url">Web page to ingest (required).</param>
        /// <param name="pipeline">Optional value to specify ingestion pipeline steps. Defaults to server configured defaults..</param>
        /// <param name="webHookUrl">Url to use for webhook callback when operation finishes or fails..</param>
        /// <param name="embeddingModel">Embedding model to use in ingestion. Optional. Default to configured default..</param>
        /// <param name="args">args.</param>
        public IngestWebPageDocumentRequest(string documentId = default(string), string index = default(string), Dictionary<string, List<string>> tags = default(Dictionary<string, List<string>>), string url = default(string), List<string> pipeline = default(List<string>), string webHookUrl = default(string), string embeddingModel = default(string), Dictionary<string, Object> args = default(Dictionary<string, Object>))
        {
            // to ensure "documentId" is required (not null)
            if (documentId == null)
            {
                throw new ArgumentNullException("documentId is a required property for IngestWebPageDocumentRequest and cannot be null");
            }
            this.DocumentId = documentId;
            // to ensure "url" is required (not null)
            if (url == null)
            {
                throw new ArgumentNullException("url is a required property for IngestWebPageDocumentRequest and cannot be null");
            }
            this.Url = url;
            this.Index = index;
            this.Tags = tags;
            this.Pipeline = pipeline;
            this.WebHookUrl = webHookUrl;
            this.EmbeddingModel = embeddingModel;
            this.Args = args;
        }

        /// <summary>
        /// Id that uniquely identifies content. Previously ingested documents with the same id will be overwritten
        /// </summary>
        /// <value>Id that uniquely identifies content. Previously ingested documents with the same id will be overwritten</value>
        [DataMember(Name = "documentId", IsRequired = true, EmitDefaultValue = true)]
        public string DocumentId { get; set; }

        /// <summary>
        /// Optional value to specify with index the document should be ingested. Defaults to &#39;default&#39;
        /// </summary>
        /// <value>Optional value to specify with index the document should be ingested. Defaults to &#39;default&#39;</value>
        [DataMember(Name = "index", EmitDefaultValue = true)]
        public string Index { get; set; }

        /// <summary>
        /// Optionally add tags to ingestion
        /// </summary>
        /// <value>Optionally add tags to ingestion</value>
        [DataMember(Name = "tags", EmitDefaultValue = true)]
        public Dictionary<string, List<string>> Tags { get; set; }

        /// <summary>
        /// Web page to ingest
        /// </summary>
        /// <value>Web page to ingest</value>
        [DataMember(Name = "url", IsRequired = true, EmitDefaultValue = true)]
        public string Url { get; set; }

        /// <summary>
        /// Optional value to specify ingestion pipeline steps. Defaults to server configured defaults.
        /// </summary>
        /// <value>Optional value to specify ingestion pipeline steps. Defaults to server configured defaults.</value>
        [DataMember(Name = "pipeline", EmitDefaultValue = true)]
        public List<string> Pipeline { get; set; }

        /// <summary>
        /// Url to use for webhook callback when operation finishes or fails.
        /// </summary>
        /// <value>Url to use for webhook callback when operation finishes or fails.</value>
        [DataMember(Name = "webHookUrl", EmitDefaultValue = true)]
        public string WebHookUrl { get; set; }

        /// <summary>
        /// Embedding model to use in ingestion. Optional. Default to configured default.
        /// </summary>
        /// <value>Embedding model to use in ingestion. Optional. Default to configured default.</value>
        [DataMember(Name = "embeddingModel", EmitDefaultValue = true)]
        public string EmbeddingModel { get; set; }

        /// <summary>
        /// Gets or Sets Args
        /// </summary>
        [DataMember(Name = "args", EmitDefaultValue = true)]
        public Dictionary<string, Object> Args { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class IngestWebPageDocumentRequest {\n");
            sb.Append("  DocumentId: ").Append(DocumentId).Append("\n");
            sb.Append("  Index: ").Append(Index).Append("\n");
            sb.Append("  Tags: ").Append(Tags).Append("\n");
            sb.Append("  Url: ").Append(Url).Append("\n");
            sb.Append("  Pipeline: ").Append(Pipeline).Append("\n");
            sb.Append("  WebHookUrl: ").Append(WebHookUrl).Append("\n");
            sb.Append("  EmbeddingModel: ").Append(EmbeddingModel).Append("\n");
            sb.Append("  Args: ").Append(Args).Append("\n");
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
            // DocumentId (string) maxLength
            if (this.DocumentId != null && this.DocumentId.Length > 255)
            {
                yield return new ValidationResult("Invalid value for DocumentId, length must be less than 255.", new [] { "DocumentId" });
            }

            // DocumentId (string) minLength
            if (this.DocumentId != null && this.DocumentId.Length < 1)
            {
                yield return new ValidationResult("Invalid value for DocumentId, length must be greater than 1.", new [] { "DocumentId" });
            }

            yield break;
        }
    }

}