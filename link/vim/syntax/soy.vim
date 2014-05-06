" Vim syntax file
" Language:	Soy Templates
" Maintainer:	Rodrigo Machado rcmachado@gmail.com
" Last Change:  Thu Apr 15 16:59:00 GMT 2010
" Filenames:    *.soy
" URL:		http://gist.github.com/gists/367358/download
"
" Based on Smarty.vim

" For version 5.x: Clear all syntax items
" For version 6.x: Quit when a syntax file was already loaded
if !exists("main_syntax")
  if version < 600
    syntax clear
  elseif exists("b:current_syntax")
    finish
  endif
  let main_syntax = 'soy'
endif

syn case ignore

runtime! syntax/html.vim
"syn cluster htmlPreproc add=smartyUnZone

syn keyword soyTagName template literal print msg namespace
syn keyword soyTagName if elseif else switch case default
syn keyword soyTagName foreach ifempty for in call param css

syn keyword smartyInFunc ne eq == != > < >= <= === ! %

" template tag
syn match soyProperty contained "private="
syn match soyProperty contained "autoescape="
" msg tag
syn match soyProperty contained "desc="
syn match soyProperty contained "meaning="
" call tag
syn match soyProperty contained "data="




syn match smartyConstant "\$smarty" 

syn match smartyDollarSign      contained "\$"
syn match smartyMaybeDollarSign contained "\([^\\]\|\\\\\)\@<=\$"

syn match smartyVariable      contained "\$\@<=\h\w*"
syn match smartyVariable      contained "\(\$\h\w*\(\.\|\->\|\[.*\]\(\.\|\->\)\)\)\@<=\w*"
syn match smartyMaybeVariable contained "\(\(^\|[^\\]\|\\\\\)\$\)\@<=\h\w*"


syn match smartyEscapedVariable contained "\\$\h\w*"

syn region smartyInBracket    matchgroup=Constant start=+\[+ end=+\]+ contains=smartyVariable contained
syn region smartyInBacktick   matchgroup=Constant start=+\`+ end=+\`+ contains=smartyVariable contained
syn region smartyStringDouble matchgroup=Constant start=+"+  end=+"+  contains=smartyMaybeVariable, smartyInBacktick, smartyMaybeDollarSign contained keepend

syn match smartyGlue "\.\|\->"


syn region smartyModifier  matchgroup=Statement start=+|+   end=+\ze:\|\>+
syn region smartyParameter matchgroup=Statement start=+:+   end=+\s\|}+ contains=smartyVariable, smartyDollarSign, smartyGlue, smartyInBracket, smartyStringDouble
syn region smartyZone     matchgroup=Statement   start="{"   end="}" contains=smartyParameter, soyProperty, smartyGlue, smartyModifier, smartyDollarSign, smartyInBracket, smartyStringDouble, smartyVariable, smartyString, smartyBlock, soyTagName, smartyConstant, smartyInFunc
syn region smartyComment  matchgroup=Comment   start="/\*\*" end="\*/"
syn region smartyComment  matchgroup=Comment   start="//" skip="\\$" end="$"

syn region  htmlString   contained start=+"+ end=+"+ contains=htmlSpecialChar,javaScriptExpression,@htmlPreproc,smartyZone
syn region  htmlString   contained start=+'+ end=+'+ contains=htmlSpecialChar,javaScriptExpression,@htmlPreproc,smartyZone
  syn region htmlLink start="<a\>\_[^>]*\<href\>" end="</a>"me=e-4 contains=@Spell,htmlTag,htmlEndTag,htmlSpecialChar,htmlPreProc,htmlComment,javaScript,@htmlPreproc,smartyZone


if version >= 508 || !exists("did_smarty_syn_inits")
  if version < 508
    let did_smarty_syn_inits = 1
    command -nargs=+ HiLink hi link <args>
  else
    command -nargs=+ HiLink hi def link <args>
  endif

  HiLink soyTagName         Function
  HiLink soyProperty        Type
  HiLink smartyComment         Comment
  HiLink smartyInFunc          Function
  HiLink smartyBlock           Constant
  HiLink smartyGlue            Statement
  HiLink smartyVariable        Identifier
  HiLink smartyDollarSign      Statement
  HiLink smartyMaybeVariable   Identifier
  HiLink smartyMaybeDollarSign Statement
  HiLink smartyStringDouble    Special
  HiLink smartyInBracket       PreProc
  HiLink smartyInBacktick      Statement
  HiLink smartyModifier        Special
  delcommand HiLink
endif 

let b:current_syntax = "soy"

if main_syntax == 'soy'
  unlet main_syntax
endif

" vim: ts=8