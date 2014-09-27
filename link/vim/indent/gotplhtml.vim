" Only load this indent file when no other was loaded.
if exists("b:did_indent")
  finish
endif

" Use HTML indentation.
runtime! indent/html.vim
