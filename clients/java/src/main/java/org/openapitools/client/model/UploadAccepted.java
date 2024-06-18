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
import java.util.Arrays;
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
 * UploadAccepted
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-06-18T14:05:27.871542961Z[Etc/UTC]", comments = "Generator version: 7.7.0-SNAPSHOT")
public class UploadAccepted {
  public static final String SERIALIZED_NAME_INDEX = "index";
  @SerializedName(SERIALIZED_NAME_INDEX)
  private String index;

  public static final String SERIALIZED_NAME_DOCUMENT_ID = "documentId";
  @SerializedName(SERIALIZED_NAME_DOCUMENT_ID)
  private String documentId;

  public static final String SERIALIZED_NAME_MESSAGE = "message";
  @SerializedName(SERIALIZED_NAME_MESSAGE)
  private String message;

  public UploadAccepted() {
  }

  public UploadAccepted index(String index) {
    this.index = index;
    return this;
  }

   /**
   * Get index
   * @return index
  **/
  @javax.annotation.Nullable
  public String getIndex() {
    return index;
  }

  public void setIndex(String index) {
    this.index = index;
  }


  public UploadAccepted documentId(String documentId) {
    this.documentId = documentId;
    return this;
  }

   /**
   * Get documentId
   * @return documentId
  **/
  @javax.annotation.Nullable
  public String getDocumentId() {
    return documentId;
  }

  public void setDocumentId(String documentId) {
    this.documentId = documentId;
  }


  public UploadAccepted message(String message) {
    this.message = message;
    return this;
  }

   /**
   * Get message
   * @return message
  **/
  @javax.annotation.Nullable
  public String getMessage() {
    return message;
  }

  public void setMessage(String message) {
    this.message = message;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UploadAccepted uploadAccepted = (UploadAccepted) o;
    return Objects.equals(this.index, uploadAccepted.index) &&
        Objects.equals(this.documentId, uploadAccepted.documentId) &&
        Objects.equals(this.message, uploadAccepted.message);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(index, documentId, message);
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
    sb.append("class UploadAccepted {\n");
    sb.append("    index: ").append(toIndentedString(index)).append("\n");
    sb.append("    documentId: ").append(toIndentedString(documentId)).append("\n");
    sb.append("    message: ").append(toIndentedString(message)).append("\n");
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
    openapiFields.add("documentId");
    openapiFields.add("message");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to UploadAccepted
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!UploadAccepted.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in UploadAccepted is not found in the empty JSON string", UploadAccepted.openapiRequiredFields.toString()));
        }
      }

      Set<Map.Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Map.Entry<String, JsonElement> entry : entries) {
        if (!UploadAccepted.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `UploadAccepted` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("index") != null && !jsonObj.get("index").isJsonNull()) && !jsonObj.get("index").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `index` to be a primitive type in the JSON string but got `%s`", jsonObj.get("index").toString()));
      }
      if ((jsonObj.get("documentId") != null && !jsonObj.get("documentId").isJsonNull()) && !jsonObj.get("documentId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `documentId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("documentId").toString()));
      }
      if ((jsonObj.get("message") != null && !jsonObj.get("message").isJsonNull()) && !jsonObj.get("message").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `message` to be a primitive type in the JSON string but got `%s`", jsonObj.get("message").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!UploadAccepted.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'UploadAccepted' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<UploadAccepted> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(UploadAccepted.class));

       return (TypeAdapter<T>) new TypeAdapter<UploadAccepted>() {
           @Override
           public void write(JsonWriter out, UploadAccepted value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public UploadAccepted read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of UploadAccepted given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of UploadAccepted
  * @throws IOException if the JSON string is invalid with respect to UploadAccepted
  */
  public static UploadAccepted fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, UploadAccepted.class);
  }

 /**
  * Convert an instance of UploadAccepted to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

