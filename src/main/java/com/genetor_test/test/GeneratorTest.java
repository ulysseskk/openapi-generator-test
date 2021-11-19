package com.genetor_test.test;

import org.apache.commons.io.IOUtils;
import org.openapitools.codegen.ClientOptInput;
import org.openapitools.codegen.DefaultGenerator;
import org.openapitools.codegen.Generator;
import org.openapitools.codegen.cmd.ListGenerators;
import org.openapitools.codegen.config.CodegenConfigurator;
import org.openapitools.codegen.languages.SpringCodegen;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.util.List;

public class GeneratorTest {
    public static void main(String[] args) {
        ListGenerators listGenerators = new ListGenerators();
        listGenerators.execute();
        CodegenConfigurator configurator = CodegenConfigurator.fromFile("generator_servers.yaml");
        configurator.setGeneratorName("spring");
        configurator.setInputSpec("generator_servers.yaml");
        configurator.setPackageName("com.ulyssesk.generators.server");
        configurator.setOutputDir("spring-files/");
        Generator generator = new DefaultGenerator(false);
        ClientOptInput clientOptInput = configurator.toClientOptInput();
        SpringCodegen springCodegen = (SpringCodegen)clientOptInput.getConfig();
        springCodegen.setBasePackage("com.ulyssesk.generators.server");
        springCodegen.setConfigPackage("com.ulyssesk.generators.server.config");
        springCodegen.setApiPackage("com.ulyssesk.generators.server.api");
        springCodegen.setModelPackage("com.ulyssesk.generators.server.model");
        clientOptInput.config(springCodegen);
        List<File> files = generator.opts(clientOptInput).generate();
        files.forEach(file -> {
            System.out.println(file.getName());
        });
    }
}
