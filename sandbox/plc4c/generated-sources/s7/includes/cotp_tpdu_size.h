/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#ifndef PLC4C_S7_READ_WRITE_COTP_TPDU_SIZE_H_
#define PLC4C_S7_READ_WRITE_COTP_TPDU_SIZE_H_

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

enum plc4c_s7_read_write_cotp_tpdu_size {
  plc4c_s7_read_write_cotp_tpdu_size_SIZE_128 = 0x07,
  plc4c_s7_read_write_cotp_tpdu_size_SIZE_256 = 0x08,
  plc4c_s7_read_write_cotp_tpdu_size_SIZE_512 = 0x09,
  plc4c_s7_read_write_cotp_tpdu_size_SIZE_1024 = 0x0a,
  plc4c_s7_read_write_cotp_tpdu_size_SIZE_2048 = 0x0b,
  plc4c_s7_read_write_cotp_tpdu_size_SIZE_4096 = 0x0c,
  plc4c_s7_read_write_cotp_tpdu_size_SIZE_8192 = 0x0d
};
typedef enum plc4c_s7_read_write_cotp_tpdu_size plc4c_s7_read_write_cotp_tpdu_size;

// Get an empty NULL-struct
plc4c_s7_read_write_cotp_tpdu_size plc4c_s7_read_write_cotp_tpdu_size_null();

int plc4c_s7_read_write_cotp_tpdu_size_num_values();


uint16_t plc4c_s7_read_write_cotp_tpdu_size_get_size_in_bytes(plc4c_s7_read_write_cotp_tpdu_size value);

#ifdef __cplusplus
}
#endif

#endif  // PLC4C_S7_READ_WRITE_COTP_TPDU_SIZE_H_
