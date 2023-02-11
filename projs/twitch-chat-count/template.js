javascript: (async() => {
  function waitForElm(selector) {
    return new Promise(resolve => {
      if (document.querySelector(selector)) {
        return resolve(document.querySelector(selector));
      }
      const observer = new MutationObserver(_ => {
        if (document.querySelector(selector)) {
          resolve(document.querySelector(selector));
          observer.disconnect();
        }
      });
      observer.observe(document.body, {
        childList: true,
        subtree: true
      });
    });
  }
  const subbed = await waitForElm(`[data-a-target="subscribed-button"]`);
  subbed.click();
  const panel = await waitForElm('div.support-panel');
  await new Promise(r => setTimeout(r, 2000));
  panel.querySelectorAll('button').forEach((btn) => btn.innerHTML.includes('Gift a specific viewer') ? btn.click() : undefined);
  const users = [{{range .}}'{{.}}',{{end}}];

  const giftables = [];
  for (let u of users) {
    const search = await waitForElm('input[placeholder="Search for a Twitch ID"]');
    const nativeInputValueSetter = Object.getOwnPropertyDescriptor(window.HTMLInputElement.prototype,'value').set;
    nativeInputValueSetter.call(search, u);
    search.dispatchEvent(new Event('input', {
      bubbles: true
    }));
    let top = await waitForElm(`button[data-user_login="${u}"]`);
    top.click();
    await new Promise(r => setTimeout(r, 1500));
    panel.querySelectorAll('button').forEach((btn) => {
      if (btn.innerHTML.includes('Gift Tier 1 Sub')) {
        btn.disabled ? undefined : giftables.push(u);
        console.log(giftables)
      }
    });
    panel.querySelectorAll('button').forEach((btn) => btn.innerHTML.includes('Change Recipient') ? btn.click() : undefined);
    if (giftables.length == 8) {
      const x = await waitForElm('button[data-a-target="modalClose"]');
      x.click();
      break;
    }
  }
  await new Promise(r => setTimeout(r, 2000));
  for (let g of giftables) {
    subbed.click();
    await new Promise(r => setTimeout(r, 2000));
    document.querySelectorAll('button').forEach((btn) => btn.innerHTML.includes('Gift a specific viewer') ? btn.click() : undefined);
    const search = await waitForElm('input[placeholder="Search for a Twitch ID"]');
    const nativeInputValueSetter = Object.getOwnPropertyDescriptor(window.HTMLInputElement.prototype,'value').set;
    nativeInputValueSetter.call(search, g);
    search.dispatchEvent(new Event('input', {
      bubbles: true
    }));
    let top = await waitForElm(`button[data-user_login="${g}"]`);
    top.click();
    await new Promise(r => setTimeout(r, 1500));
    document.querySelectorAll('button').forEach((btn) => {
      if (btn.innerHTML.includes('Gift Tier 1 Sub')) btn.click();
    });
    await new Promise(r => setTimeout(r, 1500));
    document.querySelectorAll('button').forEach((btn) => {
      if (btn.innerHTML.includes('Pay $4 .99')) btn.click();
    });
    await new Promise(r => setTimeout(r, 10000));
  }
})();
