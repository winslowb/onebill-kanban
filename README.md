# onebill-kanban 🧠

Terminal-based, Git-backed interactive Kanban board built in Go using Bubble Tea & Charm libs.

---

## 🚀 Features

- **Git-backed storage**  
  Work items stored as JSON files under `~/.onebill` (or project-root `.onebill`).

- **Kanban board UI**  
  Horizontal layout with columns: Backlog, To Do, In Progress, Test, Done.

- **Vim-style navigation & interaction**  
  - `h` / `l` → move between columns  
  - `j` / `k` → navigate tasks within a column  
  - `d` or `Enter` → view task details (description, priority, points, sprint, tags)  
  - `n` → create new task (title input)  
  - `H` / `L` → move selected task left/right status  
  - `s` → assign selected task to sprint  
  - `v` → view dashboard metrics  
  - `/` or `f` → filter board by sprint, priority, or tags  
  - `q` / `Esc` → back out of any view or quit the app  

---

## 📥 Quick Start

```bash
# Clone
git clone https://github.com/winslowb/onebill-kanban.git
cd onebill-kanban

# Ensure Go is installed (>=1.22). If not, use setup script.

# Build and run
go build -o onebill ./cmd
./onebill

✅ What You Can Do Today
Run ./onebill

Add a task: n → type title → Enter

Assign to sprint: select and press s

Move task between statuses: H / L

View Dashboard: v → q to return

Filter tasks: / or f → type and Enter → q to clear

🛠️ Future Improvements
e: inline edit task attributes (description, priority, points)

: command palette (e.g. :filter sprint-5)

Multi-field input forms (type, points, tags)

Sprint velocity charts, burn-down, etc.

Persist filter state across sessions

CI integration for task lifecycle tracking

🎯 Philosophy
Built for platform-driven workflow, onebill-kanban empowers devs to manage AI/cloud platform work entirely in the terminal. No browser needed — just 👉 go 🚀 build.

🧑‍💻 Get Involved
Submit issues / feature requests on GitHub

Share keyboard shortcuts or workflow tips!

Open-source contributions welcome 💡


