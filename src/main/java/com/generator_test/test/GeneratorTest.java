package com.generator_test.test;

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
        configurator.setGeneratorName("go");  // 生成器的类型,见readme.md
        configurator.setInputSpec("generator_servers.yaml"); // 源文件的目录
        configurator.setPackageName("test_client");
        configurator.setOutputDir("go-files/");
        Generator generator = new DefaultGenerator(false);
        List<File> files = generator.opts(configurator.toClientOptInput()).generate();
        files.forEach(file -> {
            System.out.println(file.getName());
        });
    }
}
