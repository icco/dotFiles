" vim: ts=8
" Vim syntax file
" Language:     ANTLRv3
" Maintainer:   Jörn Horstmann
" Last Change:  2006-08-02

" For version 5.x: Clear all syntax items
" For version 6.x: Quit when a syntax file was already loaded
if version < 600
  syntax clear
elseif exists("b:current_syntax")
  finish
endif

syn keyword antlrKeyword grammar lexer parser tree header members options fragment returns throws scope init

syn match  antlrCharacter '\\\(r\|n\|t\|f\|b\|"\|\'\|\\\|u\x\{4}\)' contained display

syn match antlrToken "\<[A-Z_][A-Z_0-9]\+\>"
syn match antlrRule "[a-z][a-z_0-9]+"
syn match antlrScopeVariable '$\k\+::\k\+'

syn match antlrOperator "[:;@.]"
syn match antlrOperator "[()]"
syn match antlrOperator "[{}]"
syn match antlrOperator "[\[\]]"
syn match antlrOperator "[?+*~|]"
syn match antlrOperator "[->=^]"

syn region antlrLiteral start=+'+ end=+'+ contains=antlrCharacter
syn region antlrLiteral start=+"+ end=+"+ contains=antlrCharacter

syn region antlrComment start="/\*" end="\*/"
syn match  antlrComment "//.*$"

" Define the default highlighting.
" For version 5.7 and earlier: only when not done already
" For version 5.8 and later: only when an item doesn't have highlighting yet
if version >= 508
  if version < 508
    let did_antlr_syntax_inits = 1
    command -nargs=+ HiLink hi link <args>
  else
    command -nargs=+ HiLink hi def link <args>
  endif

  HiLink antlrLiteral           Constant
  HiLink antlrCharacter         Special
  HiLink antlrComment           Comment
  HiLink antlrOperator          Operator
  HiLink antlrKeyword           Keyword
  HiLink antlrToken             PreProc
  HiLink antlrScopeVariable     Identifier
  HiLink antlrRule     	Type

  delcommand HiLink
endif

let b:current_syntax = "antlr3"
