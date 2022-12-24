<script lang="ts">
import { defineComponent, h, computed, CSSProperties } from "vue";
export default defineComponent({
  props: {
    sseStatus: { type: Boolean, required: true },
    inited: { type: Boolean, required: true }
  },
  setup(props) {
    enum COLORS {
      RED = "red",
      GREEN = "green"
    }
    enum STATUS_NAME {
      OK = "OK",
      FAIL = "Fail"
    }
    enum TAG {
      DIV = "div",
      SPAN = "span"
    }
    const BOLD = "bold";
    const initedStyle = computed<CSSProperties>(() => ({
      "color": props.inited ? COLORS.GREEN : COLORS.RED,
      "font-weight": BOLD
    }));
    const sseStatusStyle = computed<CSSProperties>(() => ({
      "color": props.sseStatus ? COLORS.GREEN : COLORS.RED,
      "font-weight": BOLD
    }));
    return () => h(TAG.DIV,
      {
        class: [
          "card",
          "mt-3"
        ]
      },
      h(TAG.DIV,
        {
          class: "card-body"
        },
        [
          h(TAG.DIV,
            [
              "sse status:: ",
              h(TAG.SPAN, {
                style: sseStatusStyle.value,
                innerHTML: props.sseStatus ? STATUS_NAME.OK : STATUS_NAME.FAIL
              })
            ]),
          h(TAG.DIV,
            [
              "auth: ",
              h(TAG.SPAN, {
                style: initedStyle.value,
                innerHTML: props.inited ? STATUS_NAME.OK : STATUS_NAME.FAIL
              })
            ])
        ]
      )
    );
  }
});
</script>
