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


package org.openapitools.client.model;

import java.util.Objects;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import org.openapitools.client.model.DocumentFilters;
import org.openapitools.jackson.nullable.JsonNullable;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import com.google.gson.TypeAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

import org.openapitools.client.JSON;

/**
 * WordCloudDocumentRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-06-18T14:05:27.871542961Z[Etc/UTC]", comments = "Generator version: 7.7.0-SNAPSHOT")
public class WordCloudDocumentRequest {
  public static final String SERIALIZED_NAME_INDEX = "index";
  @SerializedName(SERIALIZED_NAME_INDEX)
  private String index;

  public static final String SERIALIZED_NAME_FILTER = "filter";
  @SerializedName(SERIALIZED_NAME_FILTER)
  private List<DocumentFilters> filter;

  public static final String SERIALIZED_NAME_LIMIT = "limit";
  @SerializedName(SERIALIZED_NAME_LIMIT)
  private Integer limit;

  /**
   * Optional filter for specifying maximum number of entries to return. Defaults to 3
   */
  @JsonAdapter(SourceEnum.Adapter.class)
  public enum SourceEnum {
    DOCUMENT_KEY("document_key"),
    
    DOCUMENT_VALUE("document_value"),
    
    DOCUMENT_KEY_AND_VALUE("document_key_and_value"),
    
    TAG_KEY("tag_key"),
    
    TAG_VALUE("tag_value"),
    
    TAG_KER_AND_VALUE("tag_ker_and_value");

    private String value;

    SourceEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static SourceEnum fromValue(String value) {
      for (SourceEnum b : SourceEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<SourceEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final SourceEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public SourceEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return SourceEnum.fromValue(value);
      }
    }

    public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      String value = jsonElement.getAsString();
      SourceEnum.fromValue(value);
    }
  }

  public static final String SERIALIZED_NAME_SOURCE = "source";
  @SerializedName(SERIALIZED_NAME_SOURCE)
  private SourceEnum source;

  public WordCloudDocumentRequest() {
  }

  public WordCloudDocumentRequest index(String index) {
    this.index = index;
    return this;
  }

   /**
   * Optional index to specify which index to search in. Defaults to &#39;default&#39;
   * @return index
  **/
  @javax.annotation.Nullable
  public String getIndex() {
    return index;
  }

  public void setIndex(String index) {
    this.index = index;
  }


  public WordCloudDocumentRequest filter(List<DocumentFilters> filter) {
    this.filter = filter;
    return this;
  }

  public WordCloudDocumentRequest addFilterItem(DocumentFilters filterItem) {
    if (this.filter == null) {
      this.filter = new ArrayList<>();
    }
    this.filter.add(filterItem);
    return this;
  }

   /**
   * Optional filtering of document id(s) and/or tags
   * @return filter
  **/
  @javax.annotation.Nullable
  public List<DocumentFilters> getFilter() {
    return filter;
  }

  public void setFilter(List<DocumentFilters> filter) {
    this.filter = filter;
  }


  public WordCloudDocumentRequest limit(Integer limit) {
    this.limit = limit;
    return this;
  }

   /**
   * Optional filter for specifying maximum number of entries to return. Defaults to 3
   * @return limit
  **/
  @javax.annotation.Nullable
  public Integer getLimit() {
    return limit;
  }

  public void setLimit(Integer limit) {
    this.limit = limit;
  }


  public WordCloudDocumentRequest source(SourceEnum source) {
    this.source = source;
    return this;
  }

   /**
   * Optional filter for specifying maximum number of entries to return. Defaults to 3
   * @return source
  **/
  @javax.annotation.Nullable
  public SourceEnum getSource() {
    return source;
  }

  public void setSource(SourceEnum source) {
    this.source = source;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    WordCloudDocumentRequest wordCloudDocumentRequest = (WordCloudDocumentRequest) o;
    return Objects.equals(this.index, wordCloudDocumentRequest.index) &&
        Objects.equals(this.filter, wordCloudDocumentRequest.filter) &&
        Objects.equals(this.limit, wordCloudDocumentRequest.limit) &&
        Objects.equals(this.source, wordCloudDocumentRequest.source);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(index, filter, limit, source);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class WordCloudDocumentRequest {\n");
    sb.append("    index: ").append(toIndentedString(index)).append("\n");
    sb.append("    filter: ").append(toIndentedString(filter)).append("\n");
    sb.append("    limit: ").append(toIndentedString(limit)).append("\n");
    sb.append("    source: ").append(toIndentedString(source)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("index");
    openapiFields.add("filter");
    openapiFields.add("limit");
    openapiFields.add("source");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to WordCloudDocumentRequest
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!WordCloudDocumentRequest.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in WordCloudDocumentRequest is not found in the empty JSON string", WordCloudDocumentRequest.openapiRequiredFields.toString()));
        }
      }

      Set<Map.Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Map.Entry<String, JsonElement> entry : entries) {
        if (!WordCloudDocumentRequest.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `WordCloudDocumentRequest` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("index") != null && !jsonObj.get("index").isJsonNull()) && !jsonObj.get("index").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `index` to be a primitive type in the JSON string but got `%s`", jsonObj.get("index").toString()));
      }
      if (jsonObj.get("filter") != null && !jsonObj.get("filter").isJsonNull()) {
        JsonArray jsonArrayfilter = jsonObj.getAsJsonArray("filter");
        if (jsonArrayfilter != null) {
          // ensure the json data is an array
          if (!jsonObj.get("filter").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `filter` to be an array in the JSON string but got `%s`", jsonObj.get("filter").toString()));
          }

          // validate the optional field `filter` (array)
          for (int i = 0; i < jsonArrayfilter.size(); i++) {
            DocumentFilters.validateJsonElement(jsonArrayfilter.get(i));
          };
        }
      }
      if ((jsonObj.get("source") != null && !jsonObj.get("source").isJsonNull()) && !jsonObj.get("source").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `source` to be a primitive type in the JSON string but got `%s`", jsonObj.get("source").toString()));
      }
      // validate the optional field `source`
      if (jsonObj.get("source") != null && !jsonObj.get("source").isJsonNull()) {
        SourceEnum.validateJsonElement(jsonObj.get("source"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!WordCloudDocumentRequest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'WordCloudDocumentRequest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<WordCloudDocumentRequest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(WordCloudDocumentRequest.class));

       return (TypeAdapter<T>) new TypeAdapter<WordCloudDocumentRequest>() {
           @Override
           public void write(JsonWriter out, WordCloudDocumentRequest value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public WordCloudDocumentRequest read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of WordCloudDocumentRequest given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of WordCloudDocumentRequest
  * @throws IOException if the JSON string is invalid with respect to WordCloudDocumentRequest
  */
  public static WordCloudDocumentRequest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, WordCloudDocumentRequest.class);
  }

 /**
  * Convert an instance of WordCloudDocumentRequest to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

