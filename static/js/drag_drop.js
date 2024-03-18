function initDragAndDrop() {
  return {
    fileName: "No file to show",
    file: "",
    showFile: false,
    handleChange(event) {
      const file = event.target.files[0];
      if (!file) return;
      this.handleFile(file);
    },
    handleDrop(event) {
      event.preventDefault();
      const file = event.dataTransfer.files[0];
      if (!file) return;
      this.handleFile(file);
    },
    handleFile(file) {
      const reader = new FileReader();
      reader.onload = () => {
        this.file = reader.result;
        this.fileName = file.name;
        this.showFile = true;
      };
      reader.readAsDataURL(file);
    },
    handleClick() {
      this.file = "";
      this.showFile = false;
      this.fileName = "No file to show";
    },
  };
}
