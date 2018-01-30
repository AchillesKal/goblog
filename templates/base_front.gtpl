<!DOCTYPE html>
<html>
    <head>
        <meta charset=utf-8 />
        <title>{{ block "title" . }}{{end}}| Go Blog</title>
        <link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic">
        <link rel="stylesheet" href="//cdn.rawgit.com/necolas/normalize.css/master/normalize.css">
        <link rel="stylesheet" href="//cdn.rawgit.com/milligram/milligram/master/dist/milligram.min.css">
        <link rel="stylesheet" href="/public/styles.css">
    </head>
    <body>
        <header style=';' class='container'>
            <div class='logo'>
                 <a class="navigation-title" href="/">  GoBlog :B </a>
            </div>
            <ul class='navigation-list float-right'>
                <li class="navigation-item" >
                    <a class="navigation-link" href="/">Home</a>
                </li>
                <li class="navigation-item" >
                    <a class="navigation-link" href="/about">About</a>
                </li>
                <li class="navigation-item" >
                    <a class="navigation-link" href="/contact">Contact</a>
                </li>
            </ul>
        </header>
        <main class='container'>
         {{ block "content" . }}{{end}}
        </main>
        <footer class='footer'>
            <div class='container'>
                This is the footer :p
            </div>
        </footer>
    </body>
</html>