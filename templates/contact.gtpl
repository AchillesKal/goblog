{{ define "title" }}
    About
{{ end }}

{{ define "content" }}
    <h1>Contact</h1>
    <form>
        <input name='name' type="text" />
        <input name='email' type="email" />
        <textarea name='message'></textarea>
        <button>Submit</button>
    </form>
{{ end }}