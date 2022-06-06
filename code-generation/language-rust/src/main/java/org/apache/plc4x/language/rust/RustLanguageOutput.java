/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.language.rust;

import freemarker.template.Configuration;
import freemarker.template.Template;
import org.apache.commons.io.FileUtils;
import org.apache.plc4x.plugins.codegenerator.protocol.freemarker.FreemarkerLanguageOutput;
import org.apache.plc4x.plugins.codegenerator.protocol.freemarker.FreemarkerLanguageTemplateHelper;
import org.apache.plc4x.plugins.codegenerator.types.definitions.TypeDefinition;
import org.apache.plc4x.plugins.codegenerator.types.exceptions.GenerationException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.File;
import java.io.FileDescriptor;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.nio.charset.Charset;
import java.nio.file.StandardCopyOption;
import java.util.*;

public class RustLanguageOutput extends FreemarkerLanguageOutput {

    private static final Logger LOGGER = LoggerFactory.getLogger(RustLanguageOutput.class);

    private final Formatter formatter = new Formatter();

    @Override
    public String getName() {
        return "Rust";
    }

    @Override
    public Set<String> supportedOptions() {
        return Collections.singleton("package");
    }

    @Override
    public List<String> supportedOutputFlavors() {
        return Arrays.asList("read-write", "read-only", "passive");
    }

    @Override
    protected List<Template> getSpecTemplates(Configuration freemarkerConfiguration) {
        return Collections.emptyList();
    }

    @Override
    protected List<Template> getComplexTypeTemplates(Configuration freemarkerConfiguration) throws IOException {
        return List.of(freemarkerConfiguration.getTemplate("templates/rust/complex-type-template.rs.ftlh"));
    }

    @Override
    protected List<Template> getEnumTypeTemplates(Configuration freemarkerConfiguration) throws IOException {
        return Collections.singletonList(
            freemarkerConfiguration.getTemplate("templates/rust/enum-template.rs.ftlh"));
    }

    @Override
    protected List<Template> getDataIoTemplates(Configuration freemarkerConfiguration) throws IOException {
        return Collections.singletonList(
            freemarkerConfiguration.getTemplate("templates/rust/data-io-template.rs.ftlh"));
    }

    @Override
    protected FreemarkerLanguageTemplateHelper getHelper(TypeDefinition thisType, String protocolName, String flavorName, Map<String, TypeDefinition> types,
                                                         Map<String, String> options) {
        return new RustLanguageTemplateHelper(thisType, protocolName, flavorName, types, options);
    }

    @Override
    protected void postProcessTemplateOutput(File outputFile) {
/*        try {
            FileUtils.writeStringToFile(
                outputFile,
                formatter.formatSourceAndFixImports(
                    FileUtils.readFileToString(outputFile, StandardCharsets.UTF_8)
                ),
                StandardCharsets.UTF_8
            );
        } catch (IOException | FormatterException e) {
            LOGGER.error("Error formatting {}", outputFile, e);
        }
 */
    }

    @Override
    public void generate(File outputDir, String languageName, String protocolName, String outputFlavor, Map<String, TypeDefinition> types, Map<String, String> options) throws GenerationException {
        super.generate(outputDir, languageName, protocolName, outputFlavor, types, options);
        // Add post generation logic here
        try {
            InputStream inputStream = this.getClass().getClassLoader().getResourceAsStream("templates/rust/Cargo.toml");
            if (inputStream == null) {
                throw new RuntimeException("Unable to generate Cargo.toml");
            }
            java.nio.file.Files.copy(
                inputStream,
                outputDir.toPath().resolve("Cargo.toml"),
                StandardCopyOption.REPLACE_EXISTING);
            inputStream.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
        // Generate lib.rs
        try {
            StringBuilder sb = new StringBuilder();

            Collection<File> files = FileUtils.listFiles(outputDir.toPath().resolve("src").toFile(), new String[]{"rs"}, false);

            for (File file : files) {
                sb.append("mod " + file.getName().split("\\.")[0] + ";\n");
            }
            sb.append("\n");
            sb.append("#[cfg(test)]\n" +
                "mod test {\n" +
                "\n" +
                "    #[test]\n" +
                "    fn test() {\n" +
                "        println!(\"Hello world!\");\n" +
                "    }\n" +
                "}");
            FileUtils.writeStringToFile(outputDir.toPath().resolve("src").resolve("lib.rs").toFile(), sb.toString(), Charset.defaultCharset());
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}