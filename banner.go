package main

import "fmt"

func displayBanner() {
  banner :=
    `
                                                                                                    
                                                .',,:cllll:;'.                                      
                                             .:ox0KKKXXKXKKK0ko:'                                   
                                           'okKKK0OkdollloxOKKKK0o,.                                
                ...                      .oOKKOo;...  .... .'cd0KK0c.                               
                :Okl:,...                .o0Oc.  ':clldxddoc,. 'cl,.                                
               .xXKKK0Okxdlc,.             '.  ,ddolloddxkxkkd;                                     
              .oKKKKKKKOkk0KKk;              .lOc';cdddxkdoxxxk:                                    
             .dKKKKKKK0lcox0KXd.             cKc.cxdddxOc  .,:lo'                                   
            ;OKKKKKKKK0:.cx0KKl              lx.,ddddd0x.      ,,                                   
            .:dO0KXKKK0:.:k0X0;               . .:cldxkkdc.     .                                   
               .';ldOKKl.,kKXKc           .;'.      .,:ldOOd'                                       
                    .cOo.'xKKK0o,...      ,k0ko;.       .cOXl                                       
                      ;c..xKKKKKo,:oddoooolodxddl,'.     .xK;                                       
                       .  lKKKKKO:..:xOxc;'..,,;;:c;.     cx.                                       
                          .xKKKKK0l. ...,:lk0Oxxxl:okdl,   .                                        
                           .lkO00Ox, .;oddlkNXkd:..:dxO0x;                                          
                             ...... 'oddddl,cxxo'.:lxddxOKd.                                        
                                   ,olloddxl;cxl.;dddoc;:xXx..''''',,,,;;;;::::::c:.                
                                  .ld:;,',,:loc,,oxdc;;cod0Nc'oddddoooooollllllccldd;               
                                  'dddddlc::l;   ;oolddxddkNd..'''''.............',.                
                                  'dddddl::ldl'.,coc;,',;:dXd........                               
                                  .lddl:;;codc;odc;cdol:;;dKc                                       
                                   ,ododddddo..odc;';ddddkKd.                                       
                                    'ldddddd:.;ddddl;cddk0o.                                        
                                     .;odddd;,odddddddxkd,                                          
                                       .':looodddddooc:'                                            
                                           ....'.....                                               
                                                                                                    

`
  name :=
    `
888                                                        8888888b.          888      888               
888                                                        888   Y88b         888      888               
888                                                        888    888         888      888               
888  8888b.  888  888 88888b.d88b.   8888b.  88888b.       888   d88P .d88b.  888  888 88888b.   8888b.  
888     "88b  Y8bd8P' 888 "888 "88b     "88b 888 "88b      8888888P" d8P  Y8b 888 .88P 888 "88b     "88b 
888 .d888888   X88K   888  888  888 .d888888 888  888      888 T88b  88888888 888888K  888  888 .d888888 
888 888  888 .d8""8b. 888  888  888 888  888 888  888      888  T88b Y8b.     888 "88b 888  888 888  888 
888 "Y888888 888  888 888  888  888 "Y888888 888  888      888   T88b "Y8888  888  888 888  888 "Y888888 
`

  fmt.Println(hiyellow(name))
  fmt.Print(hiblue(banner))

}
