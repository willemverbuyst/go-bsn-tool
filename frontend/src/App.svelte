<script>
  import { GenerateBSN, IsValidBSN } from "../wailsjs/go/main/App.js";
  let bsnNumber = "";
  let showFeedback = false;
  let isValid = false;

  function generateBSN() {
    showFeedback = false;
    GenerateBSN(false).then((result) => (bsnNumber = result));
  }

  function isValidBSN() {
    IsValidBSN(bsnNumber.valueOf()).then((result) => {
      isValid = result;
      showFeedback = true;
    });
  }
</script>

<main>
  <header>
    <h1>BSN-tool</h1>
  </header>

  <form>
    <section>
      <button id="bsn-generator-btn" type="button" on:click={generateBSN}
        >generate BSN</button
      >
    </section>
    <section>
      <div>
        <input
          type="text"
          name="bsn-number"
          id="bsn-number"
          aria-label="bsn-number"
          autocomplete="off"
          bind:value={bsnNumber}
          on:input={() => {
            showFeedback = false;
          }}
        />
        <button id="bsn-number__copy-button" type="button">
          <i class="fa fa-copy" id="bsn-number__copy-icon"></i>
        </button>
      </div>
      {#if showFeedback && isValid}
        <p id="feedback" class="feedback__valid">this bsn is valid</p>
      {:else if showFeedback && !isValid}
        <p id="feedback" class="feedback__not-valid">this bsn is not valid</p>
      {/if}
    </section>
    <section>
      <button id="bsn-validator-btn" type="button" on:click={isValidBSN}
        >validate BSN</button
      >
    </section>
  </form>
</main>

<style>
  main {
    margin: auto;
    width: 100vw;
    height: 100vh;
    display: flex;
    align-items: center;
    flex-direction: column;
  }

  section {
    padding: 1rem;
  }

  header {
    margin-top: 5vh;
    background: #666666;
    background: -webkit-linear-gradient(to right, #666666 0%, #b8b8b8 100%);
    background: -moz-linear-gradient(to right, #666666 0%, #b8b8b8 100%);
    background: linear-gradient(to right, #666666 0%, #b8b8b8 100%);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  input {
    min-width: 35rem;
    max-width: 90vw;
    border: none;
    outline: none;
    background: none;
    font-size: 2rem;
    padding: 2rem 2rem;
    color: inherit;
    margin-bottom: 3rem;
    margin-top: 3rem;
    border-radius: 3rem;
    box-shadow:
      inset 8px 8px 8px #cbced1,
      inset -8px -8px 8px #fff;
    text-align: center;
  }

  button {
    margin-top: 1rem;
    margin-bottom: 1rem;
    border: none;
    outline: none;
    cursor: pointer;
    min-width: 35rem;
    max-width: 90vw;
    height: 6rem;
    border-radius: 3rem;
    font-size: 2rem;
    font-family: inherit;
    color: #044d8c;
    background-color: #f3f3f3;
    box-shadow:
      3px 3px 8px #b1b1b1,
      -3px -3px 8px #fff;
    transition: 0.5s;
  }

  button:hover {
    box-shadow:
      2px 2px 5px #b1b1b1,
      -2px -2px 5px #fff;
  }

  button:active {
    box-shadow:
      inset 1px 1px 3px #cbced1,
      inset -1px -1px 3px #fff;
  }

  section div {
    position: relative;
  }

  #bsn-number__copy-button {
    all: revert;
    position: absolute;
    right: 2rem;
    top: 50%;
    transform: translateY(-50%);
    border: none;
    background-color: none;
    padding: 0;
    margin: 0;
    cursor: pointer;
    display: none;
  }

  #bsn-number__copy-icon {
    font-size: 2rem;
    color: #666666;
    background-color: #f3f3f3;
    transition: 0.3s;
  }

  #bsn-number__copy-icon:hover {
    transform: scale(1.1);
  }

  form section:nth-child(2) {
    position: relative;
  }

  form section:nth-child(2) #feedback {
    width: 100%;
    position: absolute;
    bottom: -5%;
    left: 50%;
    transform: translateX(-50%);
  }

  .feedback__valid {
    color: green;
  }

  .feedback__not-valid {
    color: red;
  }
</style>
