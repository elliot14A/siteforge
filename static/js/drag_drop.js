function initDragAndDrop() {
  return {
    fileName: "no file to show",
    file: "",
    showFile: false,
    handleChange(event) {
      console.log("changed");
      const file = event.target.files[0];
      if (!file) return;
      this.handleFile(file);
    },
    handleDrop(event) {
      event.preventdefault();
      const file = event.datatransfer.files[0];
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
      this.fileName = "no file to show";
    },
  };
}

function initDragAndDropMultiple() {
  return {
    showFiles: false,
    files: [],
    handleDrop(event) {
      event.preventDefault();
      const files = event.dataTransfer.files;
      if (!files) return;
      this.handleFiles(files);
    },
    handleChange(event) {
      const files = event.target.files;
      if (!files) return;
      this.handleFiles(files);
    },
    handleFiles(files) {
      const reader = new FileReader();
      for (let i = 0; i < files.length; i++) {
        reader.onload = () => {
          this.files.push(reader.result);
          this.showFiles = true;
        };
        reader.readAsDataURL(files[i]);
      }
    },
    // @click="handleClick(index)"
    handleClick(index) {
      this.files = this.files.filter((_, i) => i !== index);
      console.log(this.files);
      if (this.files.length === 0) {
        this.showFiles = false;
      }
    },
  };
}
