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

public class ConnectedDataItem extends TypeId implements Message {

  // Accessors for discriminator values.
  public Integer getId() {
    return (int) 0x00B1;
  }

  // Properties.
  protected final int sequenceCount;
  protected final CipService service;

  // Arguments.
  protected final IntegerEncoding order;

  public ConnectedDataItem(int sequenceCount, CipService service, IntegerEncoding order) {
    super(order);
    this.sequenceCount = sequenceCount;
    this.service = service;
    this.order = order;
  }

  public int getSequenceCount() {
    return sequenceCount;
  }

  public CipService getService() {
    return service;
  }

  @Override
  protected void serializeTypeIdChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ConnectedDataItem");

    // Implicit Field (packetSize) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int packetSize = (int) ((getService().getLengthInBytes()) + (2));
    writeImplicitField(
        "packetSize",
        packetSize,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (sequenceCount)
    writeSimpleField(
        "sequenceCount",
        sequenceCount,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (service)
    writeSimpleField(
        "service",
        service,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    writeBuffer.popContext("ConnectedDataItem");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ConnectedDataItem _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (packetSize)
    lengthInBits += 16;

    // Simple field (sequenceCount)
    lengthInBits += 16;

    // Simple field (service)
    lengthInBits += service.getLengthInBits();

    return lengthInBits;
  }

  public static TypeIdBuilder staticParseTypeIdBuilder(ReadBuffer readBuffer, IntegerEncoding order)
      throws ParseException {
    readBuffer.pullContext("ConnectedDataItem");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int packetSize =
        readImplicitField(
            "packetSize",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    int sequenceCount =
        readSimpleField(
            "sequenceCount",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    CipService service =
        readSimpleField(
            "service",
            new DataReaderComplexDefault<>(
                () ->
                    CipService.staticParse(
                        readBuffer,
                        (boolean) (true),
                        (int) ((packetSize) - (2)),
                        (IntegerEncoding) (order)),
                readBuffer),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    readBuffer.closeContext("ConnectedDataItem");
    // Create the instance
    return new ConnectedDataItemBuilderImpl(sequenceCount, service, order);
  }

  public static class ConnectedDataItemBuilderImpl implements TypeId.TypeIdBuilder {
    private final int sequenceCount;
    private final CipService service;
    private final IntegerEncoding order;

    public ConnectedDataItemBuilderImpl(
        int sequenceCount, CipService service, IntegerEncoding order) {
      this.sequenceCount = sequenceCount;
      this.service = service;
      this.order = order;
    }

    public ConnectedDataItem build(IntegerEncoding order) {

      ConnectedDataItem connectedDataItem = new ConnectedDataItem(sequenceCount, service, order);
      return connectedDataItem;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ConnectedDataItem)) {
      return false;
    }
    ConnectedDataItem that = (ConnectedDataItem) o;
    return (getSequenceCount() == that.getSequenceCount())
        && (getService() == that.getService())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSequenceCount(), getService());
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
