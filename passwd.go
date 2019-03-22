--- 
vars: 
  user_name: 'javi' 
  new_password: 'bernal'

tasks:
  -name: 'Change user password' 
  become: true 
  user: 
    name: '{{ user_name }}' 
    password: '{{ new_password | password_hash("sha512") }}' 
    update_password: 'always'
