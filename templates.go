package md2min

var templContent = `<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="chrome=1">
        <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=yes">
		<title></title>
		<style> 
h1,h2,h3,h4,h5,h6,p,blockquote{margin:0;padding:0}body{font-family:"Helvetica Neue",Helvetica,"Hiragino Sans GB",Arial,sans-serif;font-size:14px;line-height:18px;color:#737373;background-color:white;margin:10px 13px 10px 13px;background-color: #F9F9F5;color: #2C3E50;}table{margin:10px 0 15px 0;border-collapse:collapse}td,th{border:1px solid #ddd;padding:3px 10px}th{padding:5px 10px}a{color:#0069d6}a:hover{color:#0050a3;text-decoration:none}a img{border:0}p{margin-bottom:9px}h1,h2,h3,h4,h5,h6{color:#404040;line-height:36px}h1{margin-bottom:18px;font-size:30px}h2{font-size:24px}h3{font-size:18px}h4{font-size:16px}h5{font-size:14px}h6{font-size:13px}hr{margin:0 0 19px;border:0;border-bottom:1px solid #ccc}blockquote{padding:13px 13px 21px 15px;margin-bottom:18px;font-family:georgia,serif;font-style:italic}blockquote:before{content:"\201C";font-size:40px;margin-left:-10px;font-family:georgia,serif;color:#eee}blockquote p{font-size:14px;font-weight:300;line-height:18px;margin-bottom:0;font-style:italic}code,pre{font-family:Monaco,Andale Mono,Courier New,monospace}code{background-color:#fee9cc;color:rgba(0,0,0,0.75);padding:1px 3px;font-size:12px;-webkit-border-radius:3px;-moz-border-radius:3px;border-radius:3px}pre{display:block;padding:5px;line-height:16px;font-size:11px;white-space:pre-wrap;word-wrap:break-word;border:1px solid #23241f;-webkit-border-radius:3px;-moz-border-radius:3px;border-radius:3px}pre code{color:#737373;font-size:11px;padding:0}sup{font-size:.83em;vertical-align:super;line-height:0}*{-webkit-print-color-adjust:exact}@media screen and (min-width:914px){body{width:854px;margin:10px auto}}@media print{body,code,pre code,h1,h2,h3,h4,h5,h6{color:black}table,pre{page-break-inside:avoid}}pre{background:#23241f}pre code{display:block;background:#23241f;padding:0}pre .tag,pre code{color:#f8f8f2}pre .keyword,pre .function,pre .literal,pre .change,pre .winutils,pre .flow,pre .lisp .title,pre .clojure .built_in,pre .nginx .title,pre .tex .special{color:#66d9ef}pre .variable,pre .params{color:#fd9720}pre .constant{color:#66d9ef}pre .title,pre .class .title,pre .css .class{color:#a6e22e}pre .attribute,pre .symbol,pre .symbol .string,pre .tag .title,pre .value,pre .css .tag{color:#f92672}pre .number,pre .preprocessor,pre .pragma,pre .regexp{color:#ae81ff}pre .tag .value,pre .string,pre .css .id,pre .subst,pre .haskell .type,pre .ruby .class .parent,pre .built_in,pre .sql .aggregate,pre .django .template_tag,pre .django .variable,pre .smalltalk .class,pre .django .filter .argument,pre .smalltalk .localvars,pre .smalltalk .array,pre .attr_selector,pre .pseudo,pre .addition,pre .stream,pre .envvar,pre .apache .tag,pre .apache .cbracket,pre .tex .command,pre .prompt{color:#e6db74}pre .comment,pre .javadoc,pre .java .annotation,pre .python .decorator,pre .template_comment,pre .pi,pre .doctype,pre .deletion,pre .shebang,pre .apache .sqbracket,pre .tex .formula{color:#a6e22d;}pre .coffeescript .javascript,pre .javascript .xml,pre .tex .formula{opacity:.5}pre .xml .javascript,pre .xml .vbscript,pre .xml .css,pre .xml .cdata{opacity:.5}
        </style>
<script type="text/javascript" src='/js/highlight.pack.js'></script>
<script>hljs.initHighlightingOnLoad();</script>
	</head>
	<body>
  <div class="container">
    <div class="nav-wrap">
      <div class="nav">{{.ListMenu}}</div>
      {{.MenuLogo}}
    </div>
  	<div class="markdown-body">{{.Content}}{{.ContentLogo}}</div>
  <div>
	</body>
</html>
`
