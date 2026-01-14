
This clean separation keeps the core torrent logic decoupled from the UI layer — easier to test and evolve.

---

## 5) Set up actions or CI/CD (optional early)

Even at MVP stage, adding simple automation signals serious engineering:

**GitHub Actions** to:
- Run lint + formatting checks
- Build front and backend
- Optionally run tests

This ensures quality from day one.

---

## 6) Plan your first issues

Turn your roadmap into actionable items:

- `feat: initialize Wails + Go project`
- `feat: UI skeleton (file picker, table list)`
- `feat: backend torrent metadata generation`
- `feat: display basic torrent stats`
- `docs: improve README with workflow guide`

Use GitHub’s issue templates if you want contributors to align with your process.

---

## 7) Describing your future fundraising trajectory

Once you’ve a working MVP and clean repo history:

- Make a **project board** (GitHub Projects) organizing features, backlog, and milestones
- Add **roadmap.md** in `docs/` detailing:
  - MVP → Beta → 1.0
  - Target audiences
  - Monetization possibilities (optional pro features, cloud seeding)
  
This helps when you pitch or share with early supporters.

---

In short:  
You’re off to a great start with `Torq` as the repo name. With a solid README, good structure, clear issues, and licensing, the project will be credible and easy to grow.  

If you want, I can help you **write the first few issue templates** or **sketch the Wails + React component scaffold** next — whatever you want to start coding first. Just say so.
::contentReference[oaicite:0]{index=0}
