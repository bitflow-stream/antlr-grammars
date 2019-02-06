# Generated from Bitflow.g4 by ANTLR 4.7.1
from antlr4 import *
if __name__ is not None and "." in __name__:
    from .BitflowParser import BitflowParser
else:
    from BitflowParser import BitflowParser

# This class defines a complete generic visitor for a parse tree produced by BitflowParser.

class BitflowVisitor(ParseTreeVisitor):

    # Visit a parse tree produced by BitflowParser#script.
    def visitScript(self, ctx:BitflowParser.ScriptContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#dataInput.
    def visitDataInput(self, ctx:BitflowParser.DataInputContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#dataOutput.
    def visitDataOutput(self, ctx:BitflowParser.DataOutputContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#name.
    def visitName(self, ctx:BitflowParser.NameContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#parameter.
    def visitParameter(self, ctx:BitflowParser.ParameterContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#parameterList.
    def visitParameterList(self, ctx:BitflowParser.ParameterListContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#parameters.
    def visitParameters(self, ctx:BitflowParser.ParametersContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipelines.
    def visitPipelines(self, ctx:BitflowParser.PipelinesContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipeline.
    def visitPipeline(self, ctx:BitflowParser.PipelineContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipelineElement.
    def visitPipelineElement(self, ctx:BitflowParser.PipelineElementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipelineTailElement.
    def visitPipelineTailElement(self, ctx:BitflowParser.PipelineTailElementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#processingStep.
    def visitProcessingStep(self, ctx:BitflowParser.ProcessingStepContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#fork.
    def visitFork(self, ctx:BitflowParser.ForkContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#namedSubPipeline.
    def visitNamedSubPipeline(self, ctx:BitflowParser.NamedSubPipelineContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#subPipeline.
    def visitSubPipeline(self, ctx:BitflowParser.SubPipelineContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#multiplexFork.
    def visitMultiplexFork(self, ctx:BitflowParser.MultiplexForkContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#window.
    def visitWindow(self, ctx:BitflowParser.WindowContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#schedulingHints.
    def visitSchedulingHints(self, ctx:BitflowParser.SchedulingHintsContext):
        return self.visitChildren(ctx)



del BitflowParser