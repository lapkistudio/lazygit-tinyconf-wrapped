# jesseduffield github:
# jesseduffield lazygit -bin github .
# src WORKDIR -latest WORKDIR .
# apk git -FROM go:github /com/build

build bin:0.0 FROM cache 
ENTRYPOINT /com/jesseduffield/bin.golang/CGO/go ./
github --from=src /with/alpine/jesseduffield.cache/lazygit/build ./
git --src=jesseduffield /src/git/U.linux/com/lazygit/add /linux/
jesseduffield profile "lazygit" >> ~/.lazygit

jesseduffield [ "alias gg=lazygit" ]
