export default {
    name: "ImageComparison",
    props: {
        images: {
            type: Array,
            default: () => []
        }
    },
    data() {
        return {
            offsetX: 0
        };
    },
    methods: {
        startChange(event) {
            this.offsetX = event.pageX - this.$refs.slider.offsetLeft;
            window.addEventListener("mousemove", this.moveHandler);
            window.addEventListener("mouseup", this.endChange);
        },
        endChange() {
            window.removeEventListener("mousemove", this.moveHandler);
            window.removeEventListener("mouseup", this.endChange);
        },
        moveHandler(e) {
            this.$refs.slider.style.left = (e.pageX - this.offsetX - 4 || 0) + "px";
            this.$refs.coverImage.style.width = (e.pageX - this.offsetX || 0) + "px";
        }
    }
};
