" Copyright (c) 2008 Google, Inc.
" Vim syntax file for Soy template system.  This primarily uses the Soy V1
" syntax, but some V2-specific syntax highlighting have been added.
" Author: Sen-po Hu <senpo@google.com>
" Documentation: http://goto/soy

if exists("b:current_syntax")
  finish
endif

syn clear

runtime! syntax/html.vim

command! -nargs=+ SoyHiLink hi def link <args>

" Soy uses the '-' character in keywords.
setlocal iskeyword+=-

" Soy comments and documentations.
" ================================

" soyComment: Matches both styles of comments (//, /* .. */)
syn region soyComment start=/\m\/\// end=/\m$/
    \ contains=soyCommentLinks containedin=htmlHead,htmlLink
syn region soyComment start=/\m\/\*/ end=/\m\*\//
    \ contains=soyDocumentationParam,soyCommentLinks
    \ containedin=htmlHead,htmlLink
" soyDocumentation*
" /**
"  * @param foo {number} desc.
"  * @param? bar {string} desc - see http://www.bar.com/
"  */
" soyDocumentationParam: Matches the '@param' and '@param?'
syn match soyDocumentationParam /\m@param\(?\)\?/ nextgroup=soyDocumentationVar skipwhite
" soyDocumentationVar: Matches the 'foo' and 'bar'
syn match soyDocumentationVar /\m\<\w\+\>/ contained nextgroup=soyDocumentationTypeWrapper skipwhite
" soyDocumentationTypeWrapper: Matches the '{' and '}'
syn region soyDocumentationTypeWrapper matchgroup=soyDocumentationParam start=/\m{/ end=/\m}/ contained contains=soyDocumentationType skipwhite keepend
" soyDocumentationType: Matches the 'number' and 'string'
syn match soyDocumentationType /\m\<.\+\>/ contained
" soyCommentLinks: Matches 'http://www.bar.com/'
syn match soyCommentLinks /\mhttps\?:\/\/\S\+/ contained

" Defining expressions (soyExpr.*)
" ================================
" Standard variables and constants

" soyExprString: Matches string literals that uses single/double quotes.
syn region soyExprString contained start=/\m"/ skip=/\m\\"/ end=/\m"/
syn region soyExprString contained start=/\m'/ skip=/\m\\'/ end=/\m'/
" soyExprVar: Matches all variables, such as: $foo?.bar?.baz[$xyz].
syn match soyExprVar contained /\m$\w\+[[:alnum:]\[\]?.$:_]*/
" soyExprConstants: Matches injected params and numeric constants.
syn match soyExprConstants contained /\m$\<ij\(\.\|\w\)*\>/ " Injected params
syn match soyExprConstants contained /\m\<\d\+\>/

" Operators and functions
" ================================

" soyExprOperator: Matches all valid soy operators.
syn keyword soyExprOperator contained is and not or
syn match soyExprOperator contained /\m\(-\|\*\|\/\|%\|+\|<\|>\|<=\|>=\|==\|!=\|?\|:\|&&\|||\)/
" soyExprFunc, soyFuncInner: Matches all (even nested) function calls.
syn keyword soyExprFunc contained isFirst isLast index length hasData keys
    \ round floor ceiling min max randomInt remainder isNonnull
    \ nextgroup=soyFuncInner skipwhite
syn region soyFuncInner start=/\m(/ end=/\m)/ matchgroup=soyFuncInner contained contains=soyExpr.*

" Commands
" ================================
" Matches all commands supported by soy.  Where appropriate, the command will
" specify the argument that may appear in valid soy code.
" Note: soyExpr.* is used to match any expression specified in the two sections
"     above.
" Note: The soyPrintCommand has to be first as the last pattern defined always
"     take priority over patterns that were defined earlier.
syn region soyPrintCommand matchgroup=soyBraces start=/\m{\(print\)\?/ end=/\m}/ contains=soyExpr.*
syn region soyParamDefCommand matchgroup=soyBraces
    \ start=/\m{@param?\?/ end=/\m}/
syn region soyInjectDefCommand matchgroup=soyBraces start=/\m{@inject/ end=/\m}/
syn region soyDelpackageCommand matchgroup=soyBraces start=/\m{delpackage/ end=/\m}/
syn region soyNamespaceCommand matchgroup=soyBraces
    \ start=/\m{namespace/ end=/\m}/
    \ contains=soyNamespaceName,soyNamespaceArg,soyExprString
syn region soyAliasCommand matchgroup=soyBraces start=/\m{alias/ end=/\m}/
    \ contains=soyNamespaceName,soyAliasSubCommand
syn region soyTemplateCommand matchgroup=soyBraces start=/\m{\(del\)\?template/ end=/\m}/ contains=soyTemplateArg,soyExprString,soyTemplateName
syn region soyCallCommand matchgroup=soyBraces start=/\m{\(del\)\?call/ end=/\m\/\?}/ contains=soyCallArg,soyExprString,soyTemplateName
syn region soyParamCommand matchgroup=soyBraces start=/\m{param/ end=/\m\/\?}/ contains=soyParamArg,soyExpr.*
syn region soyLetCommand matchgroup=soyBraces start=/\m{let/ end=/\m\/\?}/ contains=soyExpr.*
syn region soyMsgCommand matchgroup=soyBraces
    \ start=/\m{\(fallback\)\?msg/ end=/\m}/ contains=soyMsgArg,soyExprString
syn region soyCondCommand matchgroup=soyBraces start=/\m{\(if\|elseif\|else\)/ end=/\m}/ contains=soyExpr.*
syn region soyLoopCommand matchgroup=soyBraces start=/\m{for\(each\)\?/ end=/\m}/ contains=soyLoopSubCommand,soyExpr.*
syn region soySwitchCommand matchgroup=soyBraces start=/\m{switch/ end=/\m}/ contains=soyExpr.*
syn region soyCaseCommand matchgroup=soyBraces start=/\m{case/ end=/\m}/ contains=soyExpr.*
syn region soyCssCommand matchgroup=soyBraces start=/\m{css/ end=/\m}/
syn region soyPluralCommand matchgroup=soyBraces start=/\m{plural/ end=/\m}/
    \ contains=soyPluralArg,soyExpr.*
syn region soySelectCommand matchgroup=soyBraces start=/\m{select/ end=/\m}/
    \ contains=soyExpr.*
syn match soyDefaultCommand /\m{default}/
syn region soyLiteralText matchgroup=soyLiteralCommand
    \ start=/\m{literal}/ end=/\m{\/literal}/
" soyEndCommand: Matches any 'end' command, i.e. {/if}
syn match soyEndCommand /\m{\/.*}/

" Arguments that appear in soy commands
" ================================
" The arguments that is used by commands in the above section.
syn keyword soyNamespaceArg contained autoescape nextgroup=soyEqualInArg skipwhite
syn match soyNamespaceName /\m\(\.\|\w\)\+/ contained
syn keyword soyAliasSubCommand contained as
syn keyword soyTemplateArg contained name override autoescape returntype codestyle private kind nextgroup=soyEqualInArg skipwhite
syn match soyTemplateName /\m\(\.\|\w\)\+/ contained
syn keyword soyCallArg contained function data never-codestyle-stringbuilder phname nextgroup=soyEqualInArg skipwhite
syn keyword soyParamArg contained key value
syn keyword soyMsgArg contained desc hidden meaning genders
syn keyword soyLoopSubCommand contained in range
syn match soyEqualInArg /\m=/ contained
syn keyword soyPluralArg contained offset

" Non expression constants
" ================================
syn match soyCharConstants /\m{\(sp\|nil\|\\r\|\\n\|\\t\|lb\|rb\)}/

SoyHiLink soyComment Comment
SoyHiLink soyDocumentationParam Tag
SoyHiLink soyExprConstants Constant
SoyHiLink soyExprString String
SoyHiLink soyDocumentationVar Identifier
SoyHiLink soyEqualInArg Function
SoyHiLink soyFuncInner Function
SoyHiLink soyTemplateName Function
SoyHiLink soyNamespaceName Function
SoyHiLink soyExprVar Type
SoyHiLink soyNamespaceArg Type
SoyHiLink soyTemplateArg Type
SoyHiLink soyCallArg Type
SoyHiLink soyParamArg Type
SoyHiLink soyMsgArg Type
SoyHiLink soyDocumentationType Type
SoyHiLink soyPluralArg Type
SoyHiLink soyBraces Statement
SoyHiLink soyCharConstants Constant
SoyHiLink soyExprOperator Operator
SoyHiLink soyLoopSubCommand Repeat
SoyHiLink soyEndCommand Statement
SoyHiLink soyExprFunc Statement
SoyHiLink soyDefaultCommand Statement
SoyHiLink soyLiteralCommand Statement
SoyHiLink soyAliasSubCommand Statement
SoyHiLink soyCommentLinks Underlined

let b:current_syntax = "soy"
