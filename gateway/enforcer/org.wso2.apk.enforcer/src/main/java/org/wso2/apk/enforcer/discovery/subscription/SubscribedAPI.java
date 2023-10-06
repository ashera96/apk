// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/subscription/subscription.proto

package org.wso2.apk.enforcer.discovery.subscription;

/**
 * Protobuf type {@code wso2.discovery.subscription.SubscribedAPI}
 */
public final class SubscribedAPI extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:wso2.discovery.subscription.SubscribedAPI)
    SubscribedAPIOrBuilder {
private static final long serialVersionUID = 0L;
  // Use SubscribedAPI.newBuilder() to construct.
  private SubscribedAPI(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private SubscribedAPI() {
    name_ = "";
    versions_ = com.google.protobuf.LazyStringArrayList.EMPTY;
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new SubscribedAPI();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private SubscribedAPI(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();

            name_ = s;
            break;
          }
          case 18: {
            java.lang.String s = input.readStringRequireUtf8();
            if (!((mutable_bitField0_ & 0x00000001) != 0)) {
              versions_ = new com.google.protobuf.LazyStringArrayList();
              mutable_bitField0_ |= 0x00000001;
            }
            versions_.add(s);
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      if (((mutable_bitField0_ & 0x00000001) != 0)) {
        versions_ = versions_.getUnmodifiableView();
      }
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.wso2.apk.enforcer.discovery.subscription.SubscriptionProto.internal_static_wso2_discovery_subscription_SubscribedAPI_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.wso2.apk.enforcer.discovery.subscription.SubscriptionProto.internal_static_wso2_discovery_subscription_SubscribedAPI_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI.class, org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI.Builder.class);
  }

  public static final int NAME_FIELD_NUMBER = 1;
  private volatile java.lang.Object name_;
  /**
   * <code>string name = 1;</code>
   * @return The name.
   */
  @java.lang.Override
  public java.lang.String getName() {
    java.lang.Object ref = name_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      name_ = s;
      return s;
    }
  }
  /**
   * <code>string name = 1;</code>
   * @return The bytes for name.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getNameBytes() {
    java.lang.Object ref = name_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      name_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int VERSIONS_FIELD_NUMBER = 2;
  private com.google.protobuf.LazyStringList versions_;
  /**
   * <code>repeated string versions = 2;</code>
   * @return A list containing the versions.
   */
  public com.google.protobuf.ProtocolStringList
      getVersionsList() {
    return versions_;
  }
  /**
   * <code>repeated string versions = 2;</code>
   * @return The count of versions.
   */
  public int getVersionsCount() {
    return versions_.size();
  }
  /**
   * <code>repeated string versions = 2;</code>
   * @param index The index of the element to return.
   * @return The versions at the given index.
   */
  public java.lang.String getVersions(int index) {
    return versions_.get(index);
  }
  /**
   * <code>repeated string versions = 2;</code>
   * @param index The index of the value to return.
   * @return The bytes of the versions at the given index.
   */
  public com.google.protobuf.ByteString
      getVersionsBytes(int index) {
    return versions_.getByteString(index);
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!getNameBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, name_);
    }
    for (int i = 0; i < versions_.size(); i++) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, versions_.getRaw(i));
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getNameBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, name_);
    }
    {
      int dataSize = 0;
      for (int i = 0; i < versions_.size(); i++) {
        dataSize += computeStringSizeNoTag(versions_.getRaw(i));
      }
      size += dataSize;
      size += 1 * getVersionsList().size();
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI)) {
      return super.equals(obj);
    }
    org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI other = (org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI) obj;

    if (!getName()
        .equals(other.getName())) return false;
    if (!getVersionsList()
        .equals(other.getVersionsList())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + NAME_FIELD_NUMBER;
    hash = (53 * hash) + getName().hashCode();
    if (getVersionsCount() > 0) {
      hash = (37 * hash) + VERSIONS_FIELD_NUMBER;
      hash = (53 * hash) + getVersionsList().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code wso2.discovery.subscription.SubscribedAPI}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:wso2.discovery.subscription.SubscribedAPI)
      org.wso2.apk.enforcer.discovery.subscription.SubscribedAPIOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.wso2.apk.enforcer.discovery.subscription.SubscriptionProto.internal_static_wso2_discovery_subscription_SubscribedAPI_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.wso2.apk.enforcer.discovery.subscription.SubscriptionProto.internal_static_wso2_discovery_subscription_SubscribedAPI_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI.class, org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI.Builder.class);
    }

    // Construct using org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      name_ = "";

      versions_ = com.google.protobuf.LazyStringArrayList.EMPTY;
      bitField0_ = (bitField0_ & ~0x00000001);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.wso2.apk.enforcer.discovery.subscription.SubscriptionProto.internal_static_wso2_discovery_subscription_SubscribedAPI_descriptor;
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI getDefaultInstanceForType() {
      return org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI.getDefaultInstance();
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI build() {
      org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI buildPartial() {
      org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI result = new org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI(this);
      int from_bitField0_ = bitField0_;
      result.name_ = name_;
      if (((bitField0_ & 0x00000001) != 0)) {
        versions_ = versions_.getUnmodifiableView();
        bitField0_ = (bitField0_ & ~0x00000001);
      }
      result.versions_ = versions_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI) {
        return mergeFrom((org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI other) {
      if (other == org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI.getDefaultInstance()) return this;
      if (!other.getName().isEmpty()) {
        name_ = other.name_;
        onChanged();
      }
      if (!other.versions_.isEmpty()) {
        if (versions_.isEmpty()) {
          versions_ = other.versions_;
          bitField0_ = (bitField0_ & ~0x00000001);
        } else {
          ensureVersionsIsMutable();
          versions_.addAll(other.versions_);
        }
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private java.lang.Object name_ = "";
    /**
     * <code>string name = 1;</code>
     * @return The name.
     */
    public java.lang.String getName() {
      java.lang.Object ref = name_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        name_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string name = 1;</code>
     * @return The bytes for name.
     */
    public com.google.protobuf.ByteString
        getNameBytes() {
      java.lang.Object ref = name_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        name_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string name = 1;</code>
     * @param value The name to set.
     * @return This builder for chaining.
     */
    public Builder setName(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      name_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string name = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearName() {
      
      name_ = getDefaultInstance().getName();
      onChanged();
      return this;
    }
    /**
     * <code>string name = 1;</code>
     * @param value The bytes for name to set.
     * @return This builder for chaining.
     */
    public Builder setNameBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      name_ = value;
      onChanged();
      return this;
    }

    private com.google.protobuf.LazyStringList versions_ = com.google.protobuf.LazyStringArrayList.EMPTY;
    private void ensureVersionsIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        versions_ = new com.google.protobuf.LazyStringArrayList(versions_);
        bitField0_ |= 0x00000001;
       }
    }
    /**
     * <code>repeated string versions = 2;</code>
     * @return A list containing the versions.
     */
    public com.google.protobuf.ProtocolStringList
        getVersionsList() {
      return versions_.getUnmodifiableView();
    }
    /**
     * <code>repeated string versions = 2;</code>
     * @return The count of versions.
     */
    public int getVersionsCount() {
      return versions_.size();
    }
    /**
     * <code>repeated string versions = 2;</code>
     * @param index The index of the element to return.
     * @return The versions at the given index.
     */
    public java.lang.String getVersions(int index) {
      return versions_.get(index);
    }
    /**
     * <code>repeated string versions = 2;</code>
     * @param index The index of the value to return.
     * @return The bytes of the versions at the given index.
     */
    public com.google.protobuf.ByteString
        getVersionsBytes(int index) {
      return versions_.getByteString(index);
    }
    /**
     * <code>repeated string versions = 2;</code>
     * @param index The index to set the value at.
     * @param value The versions to set.
     * @return This builder for chaining.
     */
    public Builder setVersions(
        int index, java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  ensureVersionsIsMutable();
      versions_.set(index, value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string versions = 2;</code>
     * @param value The versions to add.
     * @return This builder for chaining.
     */
    public Builder addVersions(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  ensureVersionsIsMutable();
      versions_.add(value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string versions = 2;</code>
     * @param values The versions to add.
     * @return This builder for chaining.
     */
    public Builder addAllVersions(
        java.lang.Iterable<java.lang.String> values) {
      ensureVersionsIsMutable();
      com.google.protobuf.AbstractMessageLite.Builder.addAll(
          values, versions_);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string versions = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearVersions() {
      versions_ = com.google.protobuf.LazyStringArrayList.EMPTY;
      bitField0_ = (bitField0_ & ~0x00000001);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string versions = 2;</code>
     * @param value The bytes of the versions to add.
     * @return This builder for chaining.
     */
    public Builder addVersionsBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      ensureVersionsIsMutable();
      versions_.add(value);
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:wso2.discovery.subscription.SubscribedAPI)
  }

  // @@protoc_insertion_point(class_scope:wso2.discovery.subscription.SubscribedAPI)
  private static final org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI();
  }

  public static org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<SubscribedAPI>
      PARSER = new com.google.protobuf.AbstractParser<SubscribedAPI>() {
    @java.lang.Override
    public SubscribedAPI parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new SubscribedAPI(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<SubscribedAPI> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<SubscribedAPI> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.subscription.SubscribedAPI getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

