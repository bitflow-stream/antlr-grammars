// Code generated from Bitflow.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Bitflow
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBitflowListener is a complete listener for a parse tree produced by BitflowParser.
type BaseBitflowListener struct{}

var _ BitflowListener = &BaseBitflowListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBitflowListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBitflowListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBitflowListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBitflowListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseBitflowListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseBitflowListener) ExitScript(ctx *ScriptContext) {}

// EnterPipeline is called when production pipeline is entered.
func (s *BaseBitflowListener) EnterPipeline(ctx *PipelineContext) {}

// ExitPipeline is called when production pipeline is exited.
func (s *BaseBitflowListener) ExitPipeline(ctx *PipelineContext) {}

// EnterMultiInputPipeline is called when production multiInputPipeline is entered.
func (s *BaseBitflowListener) EnterMultiInputPipeline(ctx *MultiInputPipelineContext) {}

// ExitMultiInputPipeline is called when production multiInputPipeline is exited.
func (s *BaseBitflowListener) ExitMultiInputPipeline(ctx *MultiInputPipelineContext) {}

// EnterInput is called when production input is entered.
func (s *BaseBitflowListener) EnterInput(ctx *InputContext) {}

// ExitInput is called when production input is exited.
func (s *BaseBitflowListener) ExitInput(ctx *InputContext) {}

// EnterOutput is called when production output is entered.
func (s *BaseBitflowListener) EnterOutput(ctx *OutputContext) {}

// ExitOutput is called when production output is exited.
func (s *BaseBitflowListener) ExitOutput(ctx *OutputContext) {}

// EnterName is called when production name is entered.
func (s *BaseBitflowListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseBitflowListener) ExitName(ctx *NameContext) {}

// EnterNamedSubPipelineKey is called when production namedSubPipelineKey is entered.
func (s *BaseBitflowListener) EnterNamedSubPipelineKey(ctx *NamedSubPipelineKeyContext) {}

// ExitNamedSubPipelineKey is called when production namedSubPipelineKey is exited.
func (s *BaseBitflowListener) ExitNamedSubPipelineKey(ctx *NamedSubPipelineKeyContext) {}

// EnterEndpoint is called when production endpoint is entered.
func (s *BaseBitflowListener) EnterEndpoint(ctx *EndpointContext) {}

// ExitEndpoint is called when production endpoint is exited.
func (s *BaseBitflowListener) ExitEndpoint(ctx *EndpointContext) {}

// EnterVal is called when production val is entered.
func (s *BaseBitflowListener) EnterVal(ctx *ValContext) {}

// ExitVal is called when production val is exited.
func (s *BaseBitflowListener) ExitVal(ctx *ValContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseBitflowListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseBitflowListener) ExitParameter(ctx *ParameterContext) {}

// EnterTransformParameters is called when production transformParameters is entered.
func (s *BaseBitflowListener) EnterTransformParameters(ctx *TransformParametersContext) {}

// ExitTransformParameters is called when production transformParameters is exited.
func (s *BaseBitflowListener) ExitTransformParameters(ctx *TransformParametersContext) {}

// EnterIntermediateTransform is called when production intermediateTransform is entered.
func (s *BaseBitflowListener) EnterIntermediateTransform(ctx *IntermediateTransformContext) {}

// ExitIntermediateTransform is called when production intermediateTransform is exited.
func (s *BaseBitflowListener) ExitIntermediateTransform(ctx *IntermediateTransformContext) {}

// EnterTransform is called when production transform is entered.
func (s *BaseBitflowListener) EnterTransform(ctx *TransformContext) {}

// ExitTransform is called when production transform is exited.
func (s *BaseBitflowListener) ExitTransform(ctx *TransformContext) {}

// EnterFork is called when production fork is entered.
func (s *BaseBitflowListener) EnterFork(ctx *ForkContext) {}

// ExitFork is called when production fork is exited.
func (s *BaseBitflowListener) ExitFork(ctx *ForkContext) {}

// EnterNamedSubPipeline is called when production namedSubPipeline is entered.
func (s *BaseBitflowListener) EnterNamedSubPipeline(ctx *NamedSubPipelineContext) {}

// ExitNamedSubPipeline is called when production namedSubPipeline is exited.
func (s *BaseBitflowListener) ExitNamedSubPipeline(ctx *NamedSubPipelineContext) {}

// EnterSubPipeline is called when production subPipeline is entered.
func (s *BaseBitflowListener) EnterSubPipeline(ctx *SubPipelineContext) {}

// ExitSubPipeline is called when production subPipeline is exited.
func (s *BaseBitflowListener) ExitSubPipeline(ctx *SubPipelineContext) {}

// EnterMultiplexFork is called when production multiplexFork is entered.
func (s *BaseBitflowListener) EnterMultiplexFork(ctx *MultiplexForkContext) {}

// ExitMultiplexFork is called when production multiplexFork is exited.
func (s *BaseBitflowListener) ExitMultiplexFork(ctx *MultiplexForkContext) {}

// EnterMultiplexSubPipeline is called when production multiplexSubPipeline is entered.
func (s *BaseBitflowListener) EnterMultiplexSubPipeline(ctx *MultiplexSubPipelineContext) {}

// ExitMultiplexSubPipeline is called when production multiplexSubPipeline is exited.
func (s *BaseBitflowListener) ExitMultiplexSubPipeline(ctx *MultiplexSubPipelineContext) {}

// EnterWindow is called when production window is entered.
func (s *BaseBitflowListener) EnterWindow(ctx *WindowContext) {}

// ExitWindow is called when production window is exited.
func (s *BaseBitflowListener) ExitWindow(ctx *WindowContext) {}

// EnterWindowPipeline is called when production windowPipeline is entered.
func (s *BaseBitflowListener) EnterWindowPipeline(ctx *WindowPipelineContext) {}

// ExitWindowPipeline is called when production windowPipeline is exited.
func (s *BaseBitflowListener) ExitWindowPipeline(ctx *WindowPipelineContext) {}

// EnterSchedulingHints is called when production schedulingHints is entered.
func (s *BaseBitflowListener) EnterSchedulingHints(ctx *SchedulingHintsContext) {}

// ExitSchedulingHints is called when production schedulingHints is exited.
func (s *BaseBitflowListener) ExitSchedulingHints(ctx *SchedulingHintsContext) {}
