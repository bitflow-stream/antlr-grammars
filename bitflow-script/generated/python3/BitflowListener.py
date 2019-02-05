# Generated from Bitflow.g4 by ANTLR 4.7.1
from antlr4 import *
if __name__ is not None and "." in __name__:
    from .BitflowParser import BitflowParser
else:
    from BitflowParser import BitflowParser

# This class defines a complete listener for a parse tree produced by BitflowParser.
class BitflowListener(ParseTreeListener):

    # Enter a parse tree produced by BitflowParser#script.
    def enterScript(self, ctx:BitflowParser.ScriptContext):
        pass

    # Exit a parse tree produced by BitflowParser#script.
    def exitScript(self, ctx:BitflowParser.ScriptContext):
        pass


    # Enter a parse tree produced by BitflowParser#dataInput.
    def enterDataInput(self, ctx:BitflowParser.DataInputContext):
        pass

    # Exit a parse tree produced by BitflowParser#dataInput.
    def exitDataInput(self, ctx:BitflowParser.DataInputContext):
        pass


    # Enter a parse tree produced by BitflowParser#dataOutput.
    def enterDataOutput(self, ctx:BitflowParser.DataOutputContext):
        pass

    # Exit a parse tree produced by BitflowParser#dataOutput.
    def exitDataOutput(self, ctx:BitflowParser.DataOutputContext):
        pass


    # Enter a parse tree produced by BitflowParser#name.
    def enterName(self, ctx:BitflowParser.NameContext):
        pass

    # Exit a parse tree produced by BitflowParser#name.
    def exitName(self, ctx:BitflowParser.NameContext):
        pass


    # Enter a parse tree produced by BitflowParser#val.
    def enterVal(self, ctx:BitflowParser.ValContext):
        pass

    # Exit a parse tree produced by BitflowParser#val.
    def exitVal(self, ctx:BitflowParser.ValContext):
        pass


    # Enter a parse tree produced by BitflowParser#parameter.
    def enterParameter(self, ctx:BitflowParser.ParameterContext):
        pass

    # Exit a parse tree produced by BitflowParser#parameter.
    def exitParameter(self, ctx:BitflowParser.ParameterContext):
        pass


    # Enter a parse tree produced by BitflowParser#transformParameters.
    def enterTransformParameters(self, ctx:BitflowParser.TransformParametersContext):
        pass

    # Exit a parse tree produced by BitflowParser#transformParameters.
    def exitTransformParameters(self, ctx:BitflowParser.TransformParametersContext):
        pass


    # Enter a parse tree produced by BitflowParser#pipeline.
    def enterPipeline(self, ctx:BitflowParser.PipelineContext):
        pass

    # Exit a parse tree produced by BitflowParser#pipeline.
    def exitPipeline(self, ctx:BitflowParser.PipelineContext):
        pass


    # Enter a parse tree produced by BitflowParser#multiInputPipeline.
    def enterMultiInputPipeline(self, ctx:BitflowParser.MultiInputPipelineContext):
        pass

    # Exit a parse tree produced by BitflowParser#multiInputPipeline.
    def exitMultiInputPipeline(self, ctx:BitflowParser.MultiInputPipelineContext):
        pass


    # Enter a parse tree produced by BitflowParser#pipelineElement.
    def enterPipelineElement(self, ctx:BitflowParser.PipelineElementContext):
        pass

    # Exit a parse tree produced by BitflowParser#pipelineElement.
    def exitPipelineElement(self, ctx:BitflowParser.PipelineElementContext):
        pass


    # Enter a parse tree produced by BitflowParser#transform.
    def enterTransform(self, ctx:BitflowParser.TransformContext):
        pass

    # Exit a parse tree produced by BitflowParser#transform.
    def exitTransform(self, ctx:BitflowParser.TransformContext):
        pass


    # Enter a parse tree produced by BitflowParser#fork.
    def enterFork(self, ctx:BitflowParser.ForkContext):
        pass

    # Exit a parse tree produced by BitflowParser#fork.
    def exitFork(self, ctx:BitflowParser.ForkContext):
        pass


    # Enter a parse tree produced by BitflowParser#namedSubPipeline.
    def enterNamedSubPipeline(self, ctx:BitflowParser.NamedSubPipelineContext):
        pass

    # Exit a parse tree produced by BitflowParser#namedSubPipeline.
    def exitNamedSubPipeline(self, ctx:BitflowParser.NamedSubPipelineContext):
        pass


    # Enter a parse tree produced by BitflowParser#subPipeline.
    def enterSubPipeline(self, ctx:BitflowParser.SubPipelineContext):
        pass

    # Exit a parse tree produced by BitflowParser#subPipeline.
    def exitSubPipeline(self, ctx:BitflowParser.SubPipelineContext):
        pass


    # Enter a parse tree produced by BitflowParser#multiplexFork.
    def enterMultiplexFork(self, ctx:BitflowParser.MultiplexForkContext):
        pass

    # Exit a parse tree produced by BitflowParser#multiplexFork.
    def exitMultiplexFork(self, ctx:BitflowParser.MultiplexForkContext):
        pass


    # Enter a parse tree produced by BitflowParser#multiplexSubPipeline.
    def enterMultiplexSubPipeline(self, ctx:BitflowParser.MultiplexSubPipelineContext):
        pass

    # Exit a parse tree produced by BitflowParser#multiplexSubPipeline.
    def exitMultiplexSubPipeline(self, ctx:BitflowParser.MultiplexSubPipelineContext):
        pass


    # Enter a parse tree produced by BitflowParser#window.
    def enterWindow(self, ctx:BitflowParser.WindowContext):
        pass

    # Exit a parse tree produced by BitflowParser#window.
    def exitWindow(self, ctx:BitflowParser.WindowContext):
        pass


    # Enter a parse tree produced by BitflowParser#windowPipeline.
    def enterWindowPipeline(self, ctx:BitflowParser.WindowPipelineContext):
        pass

    # Exit a parse tree produced by BitflowParser#windowPipeline.
    def exitWindowPipeline(self, ctx:BitflowParser.WindowPipelineContext):
        pass


    # Enter a parse tree produced by BitflowParser#schedulingHints.
    def enterSchedulingHints(self, ctx:BitflowParser.SchedulingHintsContext):
        pass

    # Exit a parse tree produced by BitflowParser#schedulingHints.
    def exitSchedulingHints(self, ctx:BitflowParser.SchedulingHintsContext):
        pass


    # Enter a parse tree produced by BitflowParser#schedulingParameter.
    def enterSchedulingParameter(self, ctx:BitflowParser.SchedulingParameterContext):
        pass

    # Exit a parse tree produced by BitflowParser#schedulingParameter.
    def exitSchedulingParameter(self, ctx:BitflowParser.SchedulingParameterContext):
        pass


