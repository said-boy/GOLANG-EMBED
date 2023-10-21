1. golang embed
    - fitur golang embed ada di golang >= 1.16 , kurang dari itu tidak ada fitur ini

2. golang embed intinya 
    - adalah memungkinkan kita untuk dapat memasukkan sebuah file (apapun itu text,foto, html dll) kedalam program golang kita.

    - //go:embed <namafile> adalah perintah untuk menggunakan fitur golang embed ini. untuk namafile harus sama persis.

    - variabel yang akan digunakan untuk embed tidak boleh berada didalam function.

3. Image
    - gambar ketika di embed maka dia akan menghasilkan []byte. dari byte tersebut kita bisa mengubahnya kembali ke file gambar (copy) dan memberi nama berbeda.

4. Multiple File
    - jika anda ingin membaca banyak file sekaligus gunakan tipe embed.FS untuk dapat membuka satu persatu.

5. path Match
    - seperti regex.

6. hasil embed file
    - file yang sudah diembed kedalam program ketika program tersebut dibuild file tersebut akan tersimpan dan terkunci didalam build nya. berarti tidak butuh lagi file secara external. 

    - walaupun file nya dihapus atau diganti maka byte yang tersimpan kedalam variabel ketika diembed tidak akan terpengaruh ketika sudah dibuild.

    - kelebihannya adalah kita tidak perlu lagi setting file kedalam servernya nanti ketika dideploy. kita cukup bawa satu file buildnya saja.