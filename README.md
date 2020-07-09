## Requirements

- [Macports](https://www.macports.org/install.php)
- Packages `texlive`, `texlive-latex`, `texlive-latex-extra` and `texlive-fonts-extra`

## Installation

```bash
sudo port install texlive texlive-latex-extra texlive-fonts-extra
```

## Usage

Clone the repository

```bash
git clone https://github.com/melkir/resume.git
```

Run the generation

```bash
go run resume.go -lang=fr
go run resume.go -lang=en
```
