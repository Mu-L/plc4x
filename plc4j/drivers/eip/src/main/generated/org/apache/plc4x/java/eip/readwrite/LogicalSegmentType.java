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

public abstract class LogicalSegmentType implements Message {

  // Abstract accessors for discriminator values.
  public abstract Byte getLogicalSegmentType();

  // Arguments.
  protected final IntegerEncoding order;

  public LogicalSegmentType(IntegerEncoding order) {
    super();
    this.order = order;
  }

  protected abstract void serializeLogicalSegmentTypeChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("LogicalSegmentType");

    // Discriminator Field (logicalSegmentType) (Used as input to a switch field)
    writeDiscriminatorField(
        "logicalSegmentType",
        getLogicalSegmentType(),
        writeUnsignedByte(writeBuffer, 3),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Switch field (Serialize the sub-type)
    serializeLogicalSegmentTypeChild(writeBuffer);

    writeBuffer.popContext("LogicalSegmentType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    LogicalSegmentType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (logicalSegmentType)
    lengthInBits += 3;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static LogicalSegmentType staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    IntegerEncoding order;
    if (args[0] instanceof IntegerEncoding) {
      order = (IntegerEncoding) args[0];
    } else if (args[0] instanceof String) {
      order = IntegerEncoding.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type IntegerEncoding or a string which is parseable but was"
              + " "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, order);
  }

  public static LogicalSegmentType staticParse(ReadBuffer readBuffer, IntegerEncoding order)
      throws ParseException {
    readBuffer.pullContext("LogicalSegmentType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte logicalSegmentType =
        readDiscriminatorField(
            "logicalSegmentType",
            readUnsignedByte(readBuffer, 3),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    LogicalSegmentTypeBuilder builder = null;
    if (EvaluationHelper.equals(logicalSegmentType, (byte) 0x00)) {
      builder = ClassID.staticParseLogicalSegmentTypeBuilder(readBuffer, order);
    } else if (EvaluationHelper.equals(logicalSegmentType, (byte) 0x01)) {
      builder = InstanceID.staticParseLogicalSegmentTypeBuilder(readBuffer, order);
    } else if (EvaluationHelper.equals(logicalSegmentType, (byte) 0x02)) {
      builder = MemberID.staticParseLogicalSegmentTypeBuilder(readBuffer, order);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "logicalSegmentType="
              + logicalSegmentType
              + "]");
    }

    readBuffer.closeContext("LogicalSegmentType");
    // Create the instance
    LogicalSegmentType _logicalSegmentType = builder.build(order);

    return _logicalSegmentType;
  }

  public interface LogicalSegmentTypeBuilder {
    LogicalSegmentType build(IntegerEncoding order);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LogicalSegmentType)) {
      return false;
    }
    LogicalSegmentType that = (LogicalSegmentType) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
