---
mode: 'agent'
---

# 🧭 Prompt: Execute Change Plan from `plan.md`

**Goal:**
Read and interpret the research & change plan in `plan.md`, then apply the specified modifications directly to the source code files.

## 📋 Instructions

1. **Load and Parse `plan.md`**
   - Open `plan.md` and parse its sections:
     - **Analysis of User Input**
     - **Research Items**
     - **Change Plan**
     - **Next Steps**

2. **Iterate Over “Change Plan” Entries**
   For each item under the **Change Plan** section:
   - **Target Files**: Open each listed file path.
   - **Proposed Changes**:
     - Locate the lines, functions, or blocks described.
     - Apply the modifications exactly as specified.
   - **Impact & Considerations**:
     - Check for related code that might also need updating (imports, tests, docs).
     - Ensure no unintended side-effects are introduced.

3. **Validation & Testing**
   - After applying each change, run any existing automated tests or linters.
   - If a test fails, diagnose whether it’s expected (update the plan) or indicates a mis-application (fix the code).

4. **Commit & Document**
   - Stage and commit each set of related changes with a clear commit message referencing the `plan.md` section (e.g., “feat: implement X as per plan.md – Change Plan #2”).
   - If any change diverges from the plan (e.g., requires deeper refactoring), update `plan.md` to reflect the new approach.

## ⚠️ Notes

- **Accuracy First**: Don’t guess – follow the text in `plan.md` exactly.
- **One Logical Change per Commit**: Keep commits small and focused.
- **Keep `plan.md` in Sync**: If a change uncovers new research items or impacts, append those to the **Research Items** or **Next Steps** sections.
