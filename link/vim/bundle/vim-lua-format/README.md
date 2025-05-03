# vim-lua-format

Lua vim formatter supported by [LuaFormatter](https://github.com/Koihik/LuaFormatter).

## Install

Use [Vundle](https://github.com/VundleVim/Vundle.vim) to get the *vim-lua-format* plugin. After installing it you need to add the following
lines in your `.vimrc` file:

```vim
  autocmd FileType lua nnoremap <buffer> <c-k> :call LuaFormat()<cr>
  autocmd BufWrite *.lua call LuaFormat()
```

And it's done! 

Then press `<C-K>` or simply save some `*.lua` file to format the Lua code automatically.

__NOTE__ if you need to use the `LuaFormat()` function directly from command mode, you should call it explicitly as `:call LuaFormat()`

## Features

Reformats your Lua source code.

## Extension Settings

* `.lua-format`: Specifies the style config file. [Style Options](https://github.com/Koihik/LuaFormatter/wiki/Style-Config)

The `.lua-format` file must be in the source or parent directory. If there is no configuration file the default settings are used.

## Known Issues

You may have an error that claims unknown `-i` or `-si` options. This is happening because some versions of `lua-formatter` uses different flags.

So if you get any error about unknown flag, just change it to the correct flag in [flags](https://github.com/jefersonf/vim-lua-format/blob/e94e10b969bf42b76e2558d079a2765dca5baa79/autoload/lua_format.vim#L40) string variable at `lua_format#format()` function.
