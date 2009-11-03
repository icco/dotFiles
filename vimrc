""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" .vimrc -- the way it ought to be: modified from dpatierno
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

behave xterm
set nocompatible       " no compatibility with vi
filetype on            " recognize syntax by file extension
filetype indent on     " check for indent file
syntax on              " syntax highlighting
hi clear search        " do not highlight all search matches

" Don't know what these do...
set ai
set si

" meh...
"set background=light   " background light, so foreground not bold

set backspace=2        " allow <BS> to go past last insert
set gdefault           " assume :s uses /g
set ignorecase         " ignore case in search patterns
set smartcase          " searches are case-sensitive if caps used
set incsearch          " immediately highlight search matches
set noerrorbells       " no beeps on errors
set nomodeline         " prevent others from overriding this .vimrc
set number             " display line numbers
set ruler              " display row, column and % of document
set showcmd            " show partial commands in the status line
set showmatch          " show matching () {} etc.
set showmode           " show current mode

" Settings for autoindentation, comments, and what-have-you

set expandtab          " expand tabs with spaces
set tabstop=3          " <Tab> move three characters
set shiftwidth=3       " >> and << shift 3 spaces
set textwidth=79       " hard wrap at 79 characters
set modeline           " check for a modeline
set softtabstop=3      " see spaces as tabs
set scrolloff=5        " start scrolling when cursor is N lines from edge

set bg=dark

" whoa... wtf?
set nowrap             " don't soft wrap
set wrap               " linewrap

" Key Bindings, like a boss... ( cartography section )

noremap <Ins> 2<C-Y>   " <Ins> defaults like i
noremap <Del> 2<C-E>   " <Del> defaults like x

" Because we like our line numbers sometimes...
:nnoremap <C-N><C-N> :set invnumber<CR> 

" But we don't always wanna wrap
:nnoremap <C-w><C-w> :set invwrap<CR> 

" And all the cool kids need to paste
:nnoremap <C-p><C-p> :set invpaste<CR>

" Use the space key to open and close code folds
:vnoremap <space> zf<CR>
:nnoremap <space> zd<CR>

" Kinda cool
"colorscheme koehler
colorscheme darknat

" Highlights long lines
"highlight OverLength ctermbg=red ctermfg=white guibg=#592929
"match OverLength /\%81v.\+/

" Markdown
augroup mkd
  autocmd BufRead *.mkd  set ai formatoptions=tcroqn2 comments=n:>
augroup END

