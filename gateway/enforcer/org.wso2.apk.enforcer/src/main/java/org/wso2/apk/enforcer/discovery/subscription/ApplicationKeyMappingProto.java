// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/subscription/application_key_mapping.proto

package org.wso2.apk.enforcer.discovery.subscription;

public final class ApplicationKeyMappingProto {
  private ApplicationKeyMappingProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_wso2_discovery_subscription_ApplicationKeyMapping_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_wso2_discovery_subscription_ApplicationKeyMapping_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n9wso2/discovery/subscription/applicatio" +
      "n_key_mapping.proto\022\033wso2.discovery.subs" +
      "cription\"\202\001\n\025ApplicationKeyMapping\022\035\n\025ap" +
      "plicationIdentifier\030\001 \001(\t\022\017\n\007keyType\030\002 \001" +
      "(\t\022\r\n\005envId\030\003 \001(\t\022\021\n\ttimestamp\030\004 \001(\003\022\027\n\017" +
      "applicationUUID\030\005 \001(\tB\235\001\n,org.wso2.apk.e" +
      "nforcer.discovery.subscriptionB\032Applicat" +
      "ionKeyMappingProtoP\001ZOgithub.com/envoypr" +
      "oxy/go-control-plane/wso2/discovery/subs" +
      "cription;subscriptionb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_wso2_discovery_subscription_ApplicationKeyMapping_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_wso2_discovery_subscription_ApplicationKeyMapping_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_wso2_discovery_subscription_ApplicationKeyMapping_descriptor,
        new java.lang.String[] { "ApplicationIdentifier", "KeyType", "EnvId", "Timestamp", "ApplicationUUID", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
