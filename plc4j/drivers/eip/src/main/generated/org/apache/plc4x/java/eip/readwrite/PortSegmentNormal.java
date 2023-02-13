/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.eip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class PortSegmentNormal extends PortSegmentType implements Message {

  // Accessors for discriminator values.
  public Boolean getExtendedLinkAddress() {
    return (boolean) false;
  }

  // Properties.
  protected final byte port;
  protected final short linkAddress;

  // Arguments.
  protected final IntegerEncoding order;

  public PortSegmentNormal(byte port, short linkAddress, IntegerEncoding order) {
    super(order);
    this.port = port;
    this.linkAddress = linkAddress;
    this.order = order;
  }

  public byte getPort() {
    return port;
  }

  public short getLinkAddress() {
    return linkAddress;
  }

  @Override
  protected void serializePortSegmentTypeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PortSegmentNormal");

    // Simple Field (port)
    writeSimpleField("port", port, writeUnsignedByte(writeBuffer, 4));

    // Simple Field (linkAddress)
    writeSimpleField("linkAddress", linkAddress, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("PortSegmentNormal");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PortSegmentNormal _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (port)
    lengthInBits += 4;

    // Simple field (linkAddress)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static PortSegmentTypeBuilder staticParsePortSegmentTypeBuilder(
      ReadBuffer readBuffer, IntegerEncoding order) throws ParseException {
    readBuffer.pullContext("PortSegmentNormal");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte port = readSimpleField("port", readUnsignedByte(readBuffer, 4));

    short linkAddress = readSimpleField("linkAddress", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("PortSegmentNormal");
    // Create the instance
    return new PortSegmentNormalBuilderImpl(port, linkAddress, order);
  }

  public static class PortSegmentNormalBuilderImpl
      implements PortSegmentType.PortSegmentTypeBuilder {
    private final byte port;
    private final short linkAddress;
    private final IntegerEncoding order;

    public PortSegmentNormalBuilderImpl(byte port, short linkAddress, IntegerEncoding order) {
      this.port = port;
      this.linkAddress = linkAddress;
      this.order = order;
    }

    public PortSegmentNormal build(IntegerEncoding order) {

      PortSegmentNormal portSegmentNormal = new PortSegmentNormal(port, linkAddress, order);
      return portSegmentNormal;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PortSegmentNormal)) {
      return false;
    }
    PortSegmentNormal that = (PortSegmentNormal) o;
    return (getPort() == that.getPort())
        && (getLinkAddress() == that.getLinkAddress())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPort(), getLinkAddress());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
