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


    # Visit a parse tree produced by BitflowParser#val.
    def visitVal(self, ctx:BitflowParser.ValContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#parameter.
    def visitParameter(self, ctx:BitflowParser.ParameterContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#transformParameters.
    def visitTransformParameters(self, ctx:BitflowParser.TransformParametersContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipeline.
    def visitPipeline(self, ctx:BitflowParser.PipelineContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#multiInputPipeline.
    def visitMultiInputPipeline(self, ctx:BitflowParser.MultiInputPipelineContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipelineElement.
    def visitPipelineElement(self, ctx:BitflowParser.PipelineElementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#transform.
    def visitTransform(self, ctx:BitflowParser.TransformContext):
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


    # Visit a parse tree produced by BitflowParser#multiplexSubPipeline.
    def visitMultiplexSubPipeline(self, ctx:BitflowParser.MultiplexSubPipelineContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#window.
    def visitWindow(self, ctx:BitflowParser.WindowContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#windowPipeline.
    def visitWindowPipeline(self, ctx:BitflowParser.WindowPipelineContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#schedulingHints.
    def visitSchedulingHints(self, ctx:BitflowParser.SchedulingHintsContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#schedulingParameter.
    def visitSchedulingParameter(self, ctx:BitflowParser.SchedulingParameterContext):
        return self.visitChildren(ctx)



del BitflowParser