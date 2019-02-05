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


    # Visit a parse tree produced by BitflowParser#val.
    def visitVal(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#parameter.
    def visitParameter(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#transformParameters.
    def visitTransformParameters(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipeline.
    def visitPipeline(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#multiInputPipeline.
    def visitMultiInputPipeline(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#pipelineElement.
    def visitPipelineElement(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#transform.
    def visitTransform(self, ctx):
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


    # Visit a parse tree produced by BitflowParser#multiplexFork.
    def visitMultiplexFork(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#multiplexSubPipeline.
    def visitMultiplexSubPipeline(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#window.
    def visitWindow(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#windowPipeline.
    def visitWindowPipeline(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#schedulingHints.
    def visitSchedulingHints(self, ctx):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by BitflowParser#schedulingParameter.
    def visitSchedulingParameter(self, ctx):
        return self.visitChildren(ctx)


