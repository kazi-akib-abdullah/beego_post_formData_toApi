<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">

    <title>Form!</title>
  </head>
  <body>
<form class="row g-3 mt-5 needs-validation justify-content-center" action="/" method="POST" novalidate>
  <div class="col-md-8">
    <label for="first_name" class="form-label">First name</label>
    <input type="text" class="form-control" id="first_name" name="first_name" required>

    <div class="valid-feedback">
      Looks good!
    </div>
  </div>
  <div class="col-md-8">
    <label for="last_name" class="form-label">Last name</label>
    <input type="text" class="form-control" id="last_name" name="last_name" required>
    <div class="valid-feedback">
      Looks good!
    </div>
  </div>
    <div class="col-md-8">
    <label for="email" class="form-label">Email</label>
    <input type="text" class="form-control" id="email" name="email" required>
    {{if .email}}
    {{ .flash.error }}
    {{ end }}
    <div class="valid-feedback">
      Looks good!
    </div>
  </div>
    <div class="col-md-8">
    <label for="phone" class="form-label">Phone</label>
    <input type="text" class="form-control" id="phone" name="phone" required>
    {{if .phone}}
    {{ .flash.error }}
    {{ end }}
    <div class="valid-feedback">
      Looks good!
    </div>
  </div>
    <div class="col-md-8">
    <label for="password" class="form-label">Password</label>
    <input type="text" class="form-control" id="password" name="password" required>
    {{ if .password }}
    {{ .flash.error }}
    {{ end }}
    <div class="valid-feedback">
      Looks good!
    </div>
  </div>
    <div class="col-md-8">
    <label for="dob" class="form-label">Date of Birth</label>
    <input type="date" class="form-control" id="dob" name="dob" required>
    {{if .dob}}
    {{ .flash.error }}
    {{ end }}
    <div class="valid-feedback">
      Looks good!
    </div>
  </div>

  <div class="col-8">
    <button class="btn btn-primary" type="submit">Submit form</button>
      
      {{.flash.error}}

  </div>
</form>


    <!-- Option 1: Bootstrap Bundle with Popper -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p" crossorigin="anonymous"></script>

  </body>
</html>
