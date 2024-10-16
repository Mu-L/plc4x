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
package org.apache.plc4x.java.transport.serial;

import io.netty.bootstrap.Bootstrap;
import io.netty.channel.Channel;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import org.apache.plc4x.java.spi.configuration.HasConfiguration;
import org.apache.plc4x.java.spi.connection.NettyChannelFactory;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.SocketAddress;
import java.util.concurrent.Executor;

public class SerialChannelFactory extends NettyChannelFactory implements HasConfiguration<SerialTransportConfiguration> {

    private static final Logger logger = LoggerFactory.getLogger(SerialChannelFactory.class);

    private SerialTransportConfiguration configuration;

    public SerialChannelFactory(SocketAddress address) {
        super(address);
    }

    @Override
    public void setConfiguration(SerialTransportConfiguration configuration) {
        this.configuration = configuration;
    }

    @Override
    public Class<? extends Channel> getChannel() {
        return SerialChannel.class;
    }

    @Override
    public boolean isPassive() {
        return false;
    }

    @Override
    public void configureBootstrap(Bootstrap bootstrap) {
        if(configuration != null) {
            logger.info("Configuring Bootstrap with {}", configuration);
            bootstrap.option(SerialChannelOptions.BAUD_RATE, configuration.getBaudRate());
            bootstrap.option(SerialChannelOptions.DATA_BITS, configuration.getNumDataBits());
            bootstrap.option(SerialChannelOptions.STOP_BITS, configuration.getNumStopBits());
            bootstrap.option(SerialChannelOptions.PARITY, configuration.getParity());
        }
    }

    @Override
    public EventLoopGroup getEventLoopGroup() {
        return new NioEventLoopGroup(0, (Executor) null, new SerialSelectorProvider());
    }

}
