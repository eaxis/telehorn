<!DOCTYPE html>

<html>
{{ template "head.tpl"}}

<body>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <a class="navbar-brand" href="#">TeleHorn</a>

    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item active">
          <a class="nav-link" href="#">Home</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/about">About</a>
        </li>
      </ul>
    </div>
  </nav>

  <div class="container content">
    <div class="row">
      <div class="col-md-10 offset-md-1">
        <div>
          <div class="padding-bottom">
            <h3>Paste your token: </h3>
            <input class="form-control token" placeholder="153452601:AAFir5aovpw43vZJ2ReiBfgOPIb7qbd642d"/>
          </div>
          <div class="padding-bottom">
            <h3>Specify list of chats: </h3>
            <input class="form-control chats" placeholder="784178695, 785252375, 406412683..."/>
          </div>
          <div class="padding-bottom">
            <h3>Enter your message: </h3>
            <textarea class="form-control message" rows="5" placeholder="Hello fellas!"></textarea>
          </div>

          <div class="text-right padding-bottom">
            <button class="btn btn-lg btn-primary text-right send">Send messages <i class="fab fa-telegram-plane"></i></button>
          </div>
        </div>
      </div>
    </div>

    <div class="row">
      <div class="col-md-5 offset-md-7">
        <div class="alert alert-dismissible custom-alert" role="alert">
          <strong class="alert-title"></strong><br/><text class="alert-text"></text>
          <button type="button" class="close" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
      </div>
    </div>

    <div class="modal fade" tabindex="-1" role="dialog" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Are you sure?</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>

          <div class="modal-body preview">
            <img src="/static/template.png"/>
            <div class="preview-payload"></div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-primary submit">Yes, send messages</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</body>
</html>