{{ define "title" }}
    About
{{ end }}

{{ define "content" }}
    <h1>Contact</h1>
    <form method='post'>
        Name:
        <input name='name' type="text" />
        Email:
        <input name='email' type="email" />
        Message:
        <textarea name='message'></textarea>
        <button>Submit</button>
    </form>
{{ end }}