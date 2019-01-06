<!DOCTYPE html>

<html>
{{ template "head.tpl"}}

<body>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <a class="navbar-brand" href="/">TeleHorn</a>

    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item">
          <a class="nav-link" href="/">Home</a>
        </li>
        <li class="nav-item active">
          <a class="nav-link" href="/about">About</a>
        </li>
      </ul>
    </div>
  </nav>

  <div class="container content">
    <div class="row">
      <div class="col-md-10 offset-md-1 about">
        <h2>TeleHorn</h2>
        <p>
            It is simple and flexible tool to make newsletters in Telegram.<br/>
            You can use it for sending bulk messages to your users.<br/>
            See detailed docs and fork me on <a href="https://github.com/narrator69/telehorn" target="_blank">github.com/narrator69/telehorn</a>.
        </p>
      </div>
    </div>
  </div>
</body>
</html>