package com.generator_test.test;

import org.openapitools.codegen.DefaultGenerator;
import org.openapitools.codegen.Generator;
import org.openapitools.codegen.config.CodegenConfigurator;

import java.io.File;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class GeneratorTest {

    private static final Map<String, Runnable> generatorMaps = new HashMap<>();

    static {
        generatorMaps.put("go", GeneratorTest::generateGoFiles);
        generatorMaps.put("ascii", GeneratorTest::generateAsciiDoc);
        generatorMaps.put("markdown", GeneratorTest::generateMarkdownDoc);
    }

    public static void main(String[] args) {
        if (args.length < 1) {
            return;
        }
        if (args[0].equals("all")) {
            generatorMaps.forEach((k, v) -> {
                System.out.println("Executing " + k);
                v.run();
            });
        } else {
            Runnable runner = generatorMaps.get(args[0]);
            if (runner == null) {
                return;
            }
            runner.run();
        }
    }

    private static void generateGoFiles() {
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

    private static void generateAsciiDoc() {
        CodegenConfigurator configurator = CodegenConfigurator.fromFile("generator_servers.yaml");
        configurator.setGeneratorName("asciidoc");  // 生成器的类型,见readme.md
        configurator.setInputSpec("generator_servers.yaml"); // 源文件的目录
        configurator.setOutputDir("asciidoc-files/");
        Generator generator = new DefaultGenerator(false);
        List<File> files = generator.opts(configurator.toClientOptInput()).generate();
        files.forEach(file -> {
            System.out.println(file.getName());
        });
    }

    private static void generateMarkdownDoc() {
        CodegenConfigurator configurator = CodegenConfigurator.fromFile("generator_servers.yaml");
        configurator.setGeneratorName("markdown");  // 生成器的类型,见readme.md
        configurator.setInputSpec("generator_servers.yaml"); // 源文件的目录
        configurator.setOutputDir("markdown-files/");
        Generator generator = new DefaultGenerator(false);
        List<File> files = generator.opts(configurator.toClientOptInput()).generate();
        files.forEach(file -> {
            System.out.println(file.getName());
        });
    }

}
