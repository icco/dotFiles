" Vim color file
" Maintainer: Nat Welch <nat@natwelch.com>
" Last Change:	2011 April 19

" darknat -- a mod of darkblue and Dim
" [note: looks bit uglier with some terminal palettes,
" but is fine on default linux console palette.]

set background=dark
hi clear

if exists("syntax_on")
   syntax reset
endif

let colors_name = "darknat"

" Brown / Tan
"hi Statement guifg=wheat4
"hi Statement guifg=grey75 gui=bold
"hi Statement guifg=grey65 gui=bold
"hi Statement guifg=wheat4 gui=bold
"hi Statement guifg=#8B7E66 gui=bold
hi Statement guifg=#9B8E76 gui=bold

" Red
hi Constant guifg=PaleVioletRed3

" Green
"hi Identifier guifg=#00BB00
"hi Identifier guifg=#55BB55
"hi Identifier guifg=#55AA55
hi Identifier guifg=#559955

" Yellow
hi Special guifg=khaki3

" Blue
hi Comment guifg=SkyBlue3

" Purple
hi PreProc guifg=plum3

" Cyan
hi Character guifg=CadetBlue3
hi SpecialKey guifg=CadetBlue3
hi Directory guifg=SkyBlue3

" Orange
"hi Type guifg=orange4 gui=none
"hi Type guifg=orange3 gui=none
"hi Type guifg=#DD9550 gui=none
"hi Type guifg=#CD8550 gui=none
hi Type guifg=#BD7550 gui=none

hi Normal		   ctermfg=gray
hi ErrorMsg		   guifg=#ffffff guibg=#287eff ctermfg=white ctermbg=lightblue
hi Visual		   guifg=#8080ff guibg=fg gui=reverse ctermfg=lightblue ctermbg=fg cterm=reverse
hi VisualNOS	   guifg=#8080ff guibg=fg gui=reverse,underline	ctermfg=lightblue ctermbg=fg cterm=reverse,underline
hi Todo			   guifg=#d14a14 guibg=#1248d1 ctermfg=black ctermbg=gray
hi Search		   guifg=#90fff0 guibg=#2050d0 ctermfg=white ctermbg=darkblue cterm=underline term=underline
hi IncSearch	   guifg=#b0ffff guibg=#2050d0 ctermfg=darkblue ctermbg=gray

hi Title			   guifg=magenta gui=none ctermfg=magenta cterm=bold
hi WarningMsg		guifg=red ctermfg=red
hi WildMenu			guifg=yellow guibg=black ctermfg=yellow ctermbg=black cterm=none term=none
hi ModeMsg			guifg=#22cce2 ctermfg=lightblue
hi MoreMsg			ctermfg=darkgreen	ctermfg=darkgreen
hi Question			guifg=green gui=none ctermfg=green cterm=none
hi NonText			guifg=#0030ff ctermfg=darkblue

hi StatusLine	   guifg=blue guibg=darkgray gui=none ctermfg=blue ctermbg=gray term=none cterm=none
hi StatusLineNC	guifg=black guibg=darkgray gui=none ctermfg=black ctermbg=gray term=none cterm=none
hi VertSplit	   guifg=black guibg=darkgray gui=none ctermfg=black ctermbg=gray term=none cterm=none

hi Folded	      guifg=#808080 guibg=#000040 ctermfg=white ctermbg=darkgrey cterm=bold term=bold
hi FoldColumn	   guifg=#808080 guibg=#000040 ctermfg=white ctermbg=darkgrey cterm=bold term=bold
hi LineNr	      guifg=#90f020 ctermfg=cyan cterm=none

hi DiffAdd	      guibg=darkblue	ctermbg=darkblue term=none cterm=none
hi DiffChange	   guibg=darkmagenta ctermbg=magenta cterm=none
hi DiffDelete	   ctermfg=blue ctermbg=cyan gui=bold guifg=Blue guibg=DarkCyan
hi DiffText	      cterm=bold ctermbg=red gui=bold guibg=Red

hi Cursor	      guifg=black guibg=yellow ctermfg=black ctermbg=yellow
hi lCursor	      guifg=black guibg=white ctermfg=black ctermbg=white
