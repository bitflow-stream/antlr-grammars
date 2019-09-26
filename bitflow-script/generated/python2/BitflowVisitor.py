# Generated from Bitflow.g4 by ANTLR 4.7.1
from antlr4 import *

# This class defines a complete generic visitor for a parse tree produced by BitflowParser.

class BitflowVisitor(ParseTreeVisitor):

    # Visit a parse tree produced by BitflowParser#script.
    def visitScript(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#dataInput.
    def visitDataInput(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#dataOutput.
    def visitDataOutput(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#name.
    def visitName(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#parameter.
    def visitParameter(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#parameterValue.
    def visitParameterValue(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#primitiveValue.
    def visitPrimitiveValue(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#listValue.
    def visitListValue(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#mapValue.
    def visitMapValue(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#mapValueElement.
    def visitMapValueElement(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#parameterList.
    def visitParameterList(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#parameters.
    def visitParameters(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipelines.
    def visitPipelines(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipeline.
    def visitPipeline(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipelineElement.
    def visitPipelineElement(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipelineTailElement.
    def visitPipelineTailElement(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#processingStep.
    def visitProcessingStep(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#fork.
    def visitFork(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#namedSubPipeline.
    def visitNamedSubPipeline(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#subPipeline.
    def visitSubPipeline(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#batchPipeline.
    def visitBatchPipeline(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#multiplexFork.
    def visitMultiplexFork(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#batch.
    def visitBatch(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#schedulingHints.
    def visitSchedulingHints(self, ctx):
        return self.visitChildren(ctx)


