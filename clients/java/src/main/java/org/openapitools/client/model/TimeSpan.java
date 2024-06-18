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
 * TimeSpan
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-06-18T14:05:27.871542961Z[Etc/UTC]", comments = "Generator version: 7.7.0-SNAPSHOT")
public class TimeSpan {
  public static final String SERIALIZED_NAME_TICKS = "ticks";
  @SerializedName(SERIALIZED_NAME_TICKS)
  private Long ticks;

  public static final String SERIALIZED_NAME_DAYS = "days";
  @SerializedName(SERIALIZED_NAME_DAYS)
  private Integer days;

  public static final String SERIALIZED_NAME_HOURS = "hours";
  @SerializedName(SERIALIZED_NAME_HOURS)
  private Integer hours;

  public static final String SERIALIZED_NAME_MILLISECONDS = "milliseconds";
  @SerializedName(SERIALIZED_NAME_MILLISECONDS)
  private Integer milliseconds;

  public static final String SERIALIZED_NAME_MICROSECONDS = "microseconds";
  @SerializedName(SERIALIZED_NAME_MICROSECONDS)
  private Integer microseconds;

  public static final String SERIALIZED_NAME_NANOSECONDS = "nanoseconds";
  @SerializedName(SERIALIZED_NAME_NANOSECONDS)
  private Integer nanoseconds;

  public static final String SERIALIZED_NAME_MINUTES = "minutes";
  @SerializedName(SERIALIZED_NAME_MINUTES)
  private Integer minutes;

  public static final String SERIALIZED_NAME_SECONDS = "seconds";
  @SerializedName(SERIALIZED_NAME_SECONDS)
  private Integer seconds;

  public static final String SERIALIZED_NAME_TOTAL_DAYS = "totalDays";
  @SerializedName(SERIALIZED_NAME_TOTAL_DAYS)
  private Double totalDays;

  public static final String SERIALIZED_NAME_TOTAL_HOURS = "totalHours";
  @SerializedName(SERIALIZED_NAME_TOTAL_HOURS)
  private Double totalHours;

  public static final String SERIALIZED_NAME_TOTAL_MILLISECONDS = "totalMilliseconds";
  @SerializedName(SERIALIZED_NAME_TOTAL_MILLISECONDS)
  private Double totalMilliseconds;

  public static final String SERIALIZED_NAME_TOTAL_MICROSECONDS = "totalMicroseconds";
  @SerializedName(SERIALIZED_NAME_TOTAL_MICROSECONDS)
  private Double totalMicroseconds;

  public static final String SERIALIZED_NAME_TOTAL_NANOSECONDS = "totalNanoseconds";
  @SerializedName(SERIALIZED_NAME_TOTAL_NANOSECONDS)
  private Double totalNanoseconds;

  public static final String SERIALIZED_NAME_TOTAL_MINUTES = "totalMinutes";
  @SerializedName(SERIALIZED_NAME_TOTAL_MINUTES)
  private Double totalMinutes;

  public static final String SERIALIZED_NAME_TOTAL_SECONDS = "totalSeconds";
  @SerializedName(SERIALIZED_NAME_TOTAL_SECONDS)
  private Double totalSeconds;

  public TimeSpan() {
  }

  public TimeSpan(
     Integer days, 
     Integer hours, 
     Integer milliseconds, 
     Integer microseconds, 
     Integer nanoseconds, 
     Integer minutes, 
     Integer seconds, 
     Double totalDays, 
     Double totalHours, 
     Double totalMilliseconds, 
     Double totalMicroseconds, 
     Double totalNanoseconds, 
     Double totalMinutes, 
     Double totalSeconds
  ) {
    this();
    this.days = days;
    this.hours = hours;
    this.milliseconds = milliseconds;
    this.microseconds = microseconds;
    this.nanoseconds = nanoseconds;
    this.minutes = minutes;
    this.seconds = seconds;
    this.totalDays = totalDays;
    this.totalHours = totalHours;
    this.totalMilliseconds = totalMilliseconds;
    this.totalMicroseconds = totalMicroseconds;
    this.totalNanoseconds = totalNanoseconds;
    this.totalMinutes = totalMinutes;
    this.totalSeconds = totalSeconds;
  }

  public TimeSpan ticks(Long ticks) {
    this.ticks = ticks;
    return this;
  }

   /**
   * Get ticks
   * @return ticks
  **/
  @javax.annotation.Nullable
  public Long getTicks() {
    return ticks;
  }

  public void setTicks(Long ticks) {
    this.ticks = ticks;
  }


   /**
   * Get days
   * @return days
  **/
  @javax.annotation.Nullable
  public Integer getDays() {
    return days;
  }



   /**
   * Get hours
   * @return hours
  **/
  @javax.annotation.Nullable
  public Integer getHours() {
    return hours;
  }



   /**
   * Get milliseconds
   * @return milliseconds
  **/
  @javax.annotation.Nullable
  public Integer getMilliseconds() {
    return milliseconds;
  }



   /**
   * Get microseconds
   * @return microseconds
  **/
  @javax.annotation.Nullable
  public Integer getMicroseconds() {
    return microseconds;
  }



   /**
   * Get nanoseconds
   * @return nanoseconds
  **/
  @javax.annotation.Nullable
  public Integer getNanoseconds() {
    return nanoseconds;
  }



   /**
   * Get minutes
   * @return minutes
  **/
  @javax.annotation.Nullable
  public Integer getMinutes() {
    return minutes;
  }



   /**
   * Get seconds
   * @return seconds
  **/
  @javax.annotation.Nullable
  public Integer getSeconds() {
    return seconds;
  }



   /**
   * Get totalDays
   * @return totalDays
  **/
  @javax.annotation.Nullable
  public Double getTotalDays() {
    return totalDays;
  }



   /**
   * Get totalHours
   * @return totalHours
  **/
  @javax.annotation.Nullable
  public Double getTotalHours() {
    return totalHours;
  }



   /**
   * Get totalMilliseconds
   * @return totalMilliseconds
  **/
  @javax.annotation.Nullable
  public Double getTotalMilliseconds() {
    return totalMilliseconds;
  }



   /**
   * Get totalMicroseconds
   * @return totalMicroseconds
  **/
  @javax.annotation.Nullable
  public Double getTotalMicroseconds() {
    return totalMicroseconds;
  }



   /**
   * Get totalNanoseconds
   * @return totalNanoseconds
  **/
  @javax.annotation.Nullable
  public Double getTotalNanoseconds() {
    return totalNanoseconds;
  }



   /**
   * Get totalMinutes
   * @return totalMinutes
  **/
  @javax.annotation.Nullable
  public Double getTotalMinutes() {
    return totalMinutes;
  }



   /**
   * Get totalSeconds
   * @return totalSeconds
  **/
  @javax.annotation.Nullable
  public Double getTotalSeconds() {
    return totalSeconds;
  }




  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TimeSpan timeSpan = (TimeSpan) o;
    return Objects.equals(this.ticks, timeSpan.ticks) &&
        Objects.equals(this.days, timeSpan.days) &&
        Objects.equals(this.hours, timeSpan.hours) &&
        Objects.equals(this.milliseconds, timeSpan.milliseconds) &&
        Objects.equals(this.microseconds, timeSpan.microseconds) &&
        Objects.equals(this.nanoseconds, timeSpan.nanoseconds) &&
        Objects.equals(this.minutes, timeSpan.minutes) &&
        Objects.equals(this.seconds, timeSpan.seconds) &&
        Objects.equals(this.totalDays, timeSpan.totalDays) &&
        Objects.equals(this.totalHours, timeSpan.totalHours) &&
        Objects.equals(this.totalMilliseconds, timeSpan.totalMilliseconds) &&
        Objects.equals(this.totalMicroseconds, timeSpan.totalMicroseconds) &&
        Objects.equals(this.totalNanoseconds, timeSpan.totalNanoseconds) &&
        Objects.equals(this.totalMinutes, timeSpan.totalMinutes) &&
        Objects.equals(this.totalSeconds, timeSpan.totalSeconds);
  }

  @Override
  public int hashCode() {
    return Objects.hash(ticks, days, hours, milliseconds, microseconds, nanoseconds, minutes, seconds, totalDays, totalHours, totalMilliseconds, totalMicroseconds, totalNanoseconds, totalMinutes, totalSeconds);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TimeSpan {\n");
    sb.append("    ticks: ").append(toIndentedString(ticks)).append("\n");
    sb.append("    days: ").append(toIndentedString(days)).append("\n");
    sb.append("    hours: ").append(toIndentedString(hours)).append("\n");
    sb.append("    milliseconds: ").append(toIndentedString(milliseconds)).append("\n");
    sb.append("    microseconds: ").append(toIndentedString(microseconds)).append("\n");
    sb.append("    nanoseconds: ").append(toIndentedString(nanoseconds)).append("\n");
    sb.append("    minutes: ").append(toIndentedString(minutes)).append("\n");
    sb.append("    seconds: ").append(toIndentedString(seconds)).append("\n");
    sb.append("    totalDays: ").append(toIndentedString(totalDays)).append("\n");
    sb.append("    totalHours: ").append(toIndentedString(totalHours)).append("\n");
    sb.append("    totalMilliseconds: ").append(toIndentedString(totalMilliseconds)).append("\n");
    sb.append("    totalMicroseconds: ").append(toIndentedString(totalMicroseconds)).append("\n");
    sb.append("    totalNanoseconds: ").append(toIndentedString(totalNanoseconds)).append("\n");
    sb.append("    totalMinutes: ").append(toIndentedString(totalMinutes)).append("\n");
    sb.append("    totalSeconds: ").append(toIndentedString(totalSeconds)).append("\n");
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
    openapiFields.add("ticks");
    openapiFields.add("days");
    openapiFields.add("hours");
    openapiFields.add("milliseconds");
    openapiFields.add("microseconds");
    openapiFields.add("nanoseconds");
    openapiFields.add("minutes");
    openapiFields.add("seconds");
    openapiFields.add("totalDays");
    openapiFields.add("totalHours");
    openapiFields.add("totalMilliseconds");
    openapiFields.add("totalMicroseconds");
    openapiFields.add("totalNanoseconds");
    openapiFields.add("totalMinutes");
    openapiFields.add("totalSeconds");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to TimeSpan
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!TimeSpan.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in TimeSpan is not found in the empty JSON string", TimeSpan.openapiRequiredFields.toString()));
        }
      }

      Set<Map.Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Map.Entry<String, JsonElement> entry : entries) {
        if (!TimeSpan.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `TimeSpan` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!TimeSpan.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'TimeSpan' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<TimeSpan> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(TimeSpan.class));

       return (TypeAdapter<T>) new TypeAdapter<TimeSpan>() {
           @Override
           public void write(JsonWriter out, TimeSpan value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public TimeSpan read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of TimeSpan given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of TimeSpan
  * @throws IOException if the JSON string is invalid with respect to TimeSpan
  */
  public static TimeSpan fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, TimeSpan.class);
  }

 /**
  * Convert an instance of TimeSpan to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

