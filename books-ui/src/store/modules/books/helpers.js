const groupSectionsByPagesAndChapters = (sections) => {
  let groupedSections = []
  let currentChapter = null
  for (const section of sections) {
    if (section.level === 'chapter') {
      currentChapter = section
      currentChapter.items = []
      groupedSections = [...groupedSections, currentChapter]
    } else {
      currentChapter.items.push(section)
    }
  }
  return groupedSections
}

export { groupSectionsByPagesAndChapters }