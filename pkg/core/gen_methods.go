package core

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/MOXHATKA/gote/pkg/types"
)

// URL - адресс Telegram Bot API
const URL = "https://api.telegram.org/bot"

type tgResponse[T any] struct {
	Ok          bool                      `json:"ok"`
	Result      T                         `json:"result"`
	Description string                    `json:"description,omitempty"`
	ErrorCode   int                       `json:"error_code,omitempty"`
	Parameters  *types.ResponseParameters `json:"parameters,omitempty"`
}

// Use this method to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
//
// Notes1. This method will not work if an outgoing webhook is set up.2. In order to avoid getting duplicate updates, recalculate offset after each server response.
//
// https://core.telegram.org/bots/api#getupdates
func (bot *Bot) GetUpdates(ctx context.Context, param types.GetUpdates) ([]types.Update, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetUpdates"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[[]types.Update]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request (a request with response HTTP status code different from 2XY), we will repeat the request and give up after a reasonable amount of attempts. Returns True on success.
//
// Notes1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.3. Ports currently supported for webhooks: 443, 80, 88, 8443.If you&#39;re having any trouble setting up webhooks, please check out this amazing guide to webhooks.
//
// https://core.telegram.org/bots/api#setwebhook
func (bot *Bot) SetWebhook(ctx context.Context, param types.SetWebhook) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetWebhook"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletewebhook
func (bot *Bot) DeleteWebhook(ctx context.Context, param types.DeleteWebhook) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteWebhook"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
// 
// https://core.telegram.org/bots/api#getwebhookinfo
func (bot *Bot) GetWebhookInfo(ctx context.Context, param types.GetWebhookInfo) (*types.WebhookInfo, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetWebhookInfo"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.WebhookInfo]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// A simple method for testing your bot&#39;s authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
// 
// https://core.telegram.org/bots/api#getme
func (bot *Bot) GetMe(ctx context.Context, param types.GetMe) (*types.User, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetMe"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.User]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
// 
// https://core.telegram.org/bots/api#logout
func (bot *Bot) LogOut(ctx context.Context, param types.LogOut) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/LogOut"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn&#39;t launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
// 
// https://core.telegram.org/bots/api#close
func (bot *Bot) Close(ctx context.Context, param types.Close) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/Close"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send text messages. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendmessage
func (bot *Bot) SendMessage(ctx context.Context, param types.SendMessage) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendMessage"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to forward messages of any kind. Service messages and messages with protected content can&#39;t be forwarded. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#forwardmessage
func (bot *Bot) ForwardMessage(ctx context.Context, param types.ForwardMessage) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/ForwardMessage"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to forward multiple messages of any kind. If some of the specified messages can&#39;t be found or forwarded, they are skipped. Service messages and messages with protected content can&#39;t be forwarded. Album grouping is kept for forwarded messages. On success, an array of MessageId of the sent messages is returned.
// 
// https://core.telegram.org/bots/api#forwardmessages
func (bot *Bot) ForwardMessages(ctx context.Context, param types.ForwardMessages) (*types.MessageId, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/ForwardMessages"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.MessageId]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to copy messages of any kind. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can&#39;t be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn&#39;t have a link to the original message. Returns the MessageId of the sent message on success.
// 
// https://core.telegram.org/bots/api#copymessage
func (bot *Bot) CopyMessage(ctx context.Context, param types.CopyMessage) (*types.MessageId, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/CopyMessage"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.MessageId]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to copy messages of any kind. If some of the specified messages can&#39;t be found or copied, they are skipped. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can&#39;t be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessages, but the copied messages don&#39;t have a link to the original message. Album grouping is kept for copied messages. On success, an array of MessageId of the sent messages is returned.
// 
// https://core.telegram.org/bots/api#copymessages
func (bot *Bot) CopyMessages(ctx context.Context, param types.CopyMessages) (*types.MessageId, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/CopyMessages"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.MessageId]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send photos. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendphoto
func (bot *Bot) SendPhoto(ctx context.Context, param types.SendPhoto) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendPhoto"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// 
// https://core.telegram.org/bots/api#sendaudio
func (bot *Bot) SendAudio(ctx context.Context, param types.SendAudio) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendAudio"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
// 
// https://core.telegram.org/bots/api#senddocument
func (bot *Bot) SendDocument(ctx context.Context, param types.SendDocument) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendDocument"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
// 
// https://core.telegram.org/bots/api#sendvideo
func (bot *Bot) SendVideo(ctx context.Context, param types.SendVideo) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendVideo"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
// 
// https://core.telegram.org/bots/api#sendanimation
func (bot *Bot) SendAnimation(ctx context.Context, param types.SendAnimation) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendAnimation"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
// 
// https://core.telegram.org/bots/api#sendvoice
func (bot *Bot) SendVoice(ctx context.Context, param types.SendVoice) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendVoice"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendvideonote
func (bot *Bot) SendVideoNote(ctx context.Context, param types.SendVideoNote) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendVideoNote"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send paid media. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendpaidmedia
func (bot *Bot) SendPaidMedia(ctx context.Context, param types.SendPaidMedia) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendPaidMedia"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Message objects that were sent is returned.
// 
// https://core.telegram.org/bots/api#sendmediagroup
func (bot *Bot) SendMediaGroup(ctx context.Context, param types.SendMediaGroup) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendMediaGroup"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send point on the map. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendlocation
func (bot *Bot) SendLocation(ctx context.Context, param types.SendLocation) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendLocation"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send information about a venue. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendvenue
func (bot *Bot) SendVenue(ctx context.Context, param types.SendVenue) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendVenue"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send phone contacts. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendcontact
func (bot *Bot) SendContact(ctx context.Context, param types.SendContact) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendContact"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send a native poll. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendpoll
func (bot *Bot) SendPoll(ctx context.Context, param types.SendPoll) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendPoll"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send a checklist on behalf of a connected business account. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendchecklist
func (bot *Bot) SendChecklist(ctx context.Context, param types.SendChecklist) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendChecklist"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#senddice
func (bot *Bot) SendDice(ctx context.Context, param types.SendDice) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendDice"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method when you need to tell the user that something is happening on the bot&#39;s side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
//
// Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.
//
// https://core.telegram.org/bots/api#sendchataction
func (bot *Bot) SendChatAction(ctx context.Context, param types.SendChatAction) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SendChatAction"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the chosen reactions on a message. Service messages of some types can&#39;t be reacted to. Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel. Bots can&#39;t use paid reactions. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmessagereaction
func (bot *Bot) SetMessageReaction(ctx context.Context, param types.SetMessageReaction) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetMessageReaction"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
// 
// https://core.telegram.org/bots/api#getuserprofilephotos
func (bot *Bot) GetUserProfilePhotos(ctx context.Context, param types.GetUserProfilePhotos) (*types.UserProfilePhotos, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetUserProfilePhotos"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.UserProfilePhotos]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the Mini App method requestEmojiStatusAccess. Returns True on success.
// 
// https://core.telegram.org/bots/api#setuseremojistatus
func (bot *Bot) SetUserEmojiStatus(ctx context.Context, param types.SetUserEmojiStatus) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetUserEmojiStatus"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot&lt;token&gt;/&lt;file_path&gt;, where &lt;file_path&gt; is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
// 
// https://core.telegram.org/bots/api#getfile
func (bot *Bot) GetFile(ctx context.Context, param types.GetFile) (*types.File, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetFile"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.File]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#banchatmember
func (bot *Bot) BanChatMember(ctx context.Context, param types.BanChatMember) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/BanChatMember"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don&#39;t want this, use the parameter only_if_banned. Returns True on success.
// 
// https://core.telegram.org/bots/api#unbanchatmember
func (bot *Bot) UnbanChatMember(ctx context.Context, param types.UnbanChatMember) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/UnbanChatMember"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
// 
// https://core.telegram.org/bots/api#restrictchatmember
func (bot *Bot) RestrictChatMember(ctx context.Context, param types.RestrictChatMember) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/RestrictChatMember"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
// 
// https://core.telegram.org/bots/api#promotechatmember
func (bot *Bot) PromoteChatMember(ctx context.Context, param types.PromoteChatMember) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/PromoteChatMember"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (bot *Bot) SetChatAdministratorCustomTitle(ctx context.Context, param types.SetChatAdministratorCustomTitle) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetChatAdministratorCustomTitle"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won&#39;t be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#banchatsenderchat
func (bot *Bot) BanChatSenderChat(ctx context.Context, param types.BanChatSenderChat) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/BanChatSenderChat"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#unbanchatsenderchat
func (bot *Bot) UnbanChatSenderChat(ctx context.Context, param types.UnbanChatSenderChat) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/UnbanChatSenderChat"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatpermissions
func (bot *Bot) SetChatPermissions(ctx context.Context, param types.SetChatPermissions) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetChatPermissions"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots can&#39;t use invite links generated by other administrators. If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink or by calling the getChat method. If your bot needs to generate a new primary invite link replacing its previous one, use exportChatInviteLink again.
//
// https://core.telegram.org/bots/api#exportchatinvitelink
func (bot *Bot) ExportChatInviteLink(ctx context.Context, param types.ExportChatInviteLink) (string, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	
	url := URL + bot.token + "/ExportChatInviteLink"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return "", err
	}
	
	var result tgResponse[string]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
// 
// https://core.telegram.org/bots/api#createchatinvitelink
func (bot *Bot) CreateChatInviteLink(ctx context.Context, param types.CreateChatInviteLink) (*types.ChatInviteLink, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/CreateChatInviteLink"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.ChatInviteLink]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
// 
// https://core.telegram.org/bots/api#editchatinvitelink
func (bot *Bot) EditChatInviteLink(ctx context.Context, param types.EditChatInviteLink) (*types.ChatInviteLink, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/EditChatInviteLink"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.ChatInviteLink]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to create a subscription invite link for a channel chat. The bot must have the can_invite_users administrator rights. The link can be edited using the method editChatSubscriptionInviteLink or revoked using the method revokeChatInviteLink. Returns the new invite link as a ChatInviteLink object.
// 
// https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
func (bot *Bot) CreateChatSubscriptionInviteLink(ctx context.Context, param types.CreateChatSubscriptionInviteLink) (*types.ChatInviteLink, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/CreateChatSubscriptionInviteLink"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.ChatInviteLink]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit a subscription invite link created by the bot. The bot must have the can_invite_users administrator rights. Returns the edited invite link as a ChatInviteLink object.
// 
// https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
func (bot *Bot) EditChatSubscriptionInviteLink(ctx context.Context, param types.EditChatSubscriptionInviteLink) (*types.ChatInviteLink, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/EditChatSubscriptionInviteLink"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.ChatInviteLink]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
// 
// https://core.telegram.org/bots/api#revokechatinvitelink
func (bot *Bot) RevokeChatInviteLink(ctx context.Context, param types.RevokeChatInviteLink) (*types.ChatInviteLink, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/RevokeChatInviteLink"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.ChatInviteLink]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
// 
// https://core.telegram.org/bots/api#approvechatjoinrequest
func (bot *Bot) ApproveChatJoinRequest(ctx context.Context, param types.ApproveChatJoinRequest) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/ApproveChatJoinRequest"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
// 
// https://core.telegram.org/bots/api#declinechatjoinrequest
func (bot *Bot) DeclineChatJoinRequest(ctx context.Context, param types.DeclineChatJoinRequest) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeclineChatJoinRequest"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set a new profile photo for the chat. Photos can&#39;t be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatphoto
func (bot *Bot) SetChatPhoto(ctx context.Context, param types.SetChatPhoto) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetChatPhoto"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a chat photo. Photos can&#39;t be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletechatphoto
func (bot *Bot) DeleteChatPhoto(ctx context.Context, param types.DeleteChatPhoto) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteChatPhoto"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the title of a chat. Titles can&#39;t be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchattitle
func (bot *Bot) SetChatTitle(ctx context.Context, param types.SetChatTitle) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetChatTitle"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatdescription
func (bot *Bot) SetChatDescription(ctx context.Context, param types.SetChatDescription) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetChatDescription"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to add a message to the list of pinned messages in a chat. In private chats and channel direct messages chats, all non-service messages can be pinned. Conversely, the bot must be an administrator with the &#39;can_pin_messages&#39; right or the &#39;can_edit_messages&#39; right to pin messages in groups and channels respectively. Returns True on success.
// 
// https://core.telegram.org/bots/api#pinchatmessage
func (bot *Bot) PinChatMessage(ctx context.Context, param types.PinChatMessage) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/PinChatMessage"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to remove a message from the list of pinned messages in a chat. In private chats and channel direct messages chats, all messages can be unpinned. Conversely, the bot must be an administrator with the &#39;can_pin_messages&#39; right or the &#39;can_edit_messages&#39; right to unpin messages in groups and channels respectively. Returns True on success.
// 
// https://core.telegram.org/bots/api#unpinchatmessage
func (bot *Bot) UnpinChatMessage(ctx context.Context, param types.UnpinChatMessage) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/UnpinChatMessage"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to clear the list of pinned messages in a chat. In private chats and channel direct messages chats, no additional rights are required to unpin all pinned messages. Conversely, the bot must be an administrator with the &#39;can_pin_messages&#39; right or the &#39;can_edit_messages&#39; right to unpin all pinned messages in groups and channels respectively. Returns True on success.
// 
// https://core.telegram.org/bots/api#unpinallchatmessages
func (bot *Bot) UnpinAllChatMessages(ctx context.Context, param types.UnpinAllChatMessages) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/UnpinAllChatMessages"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
// 
// https://core.telegram.org/bots/api#leavechat
func (bot *Bot) LeaveChat(ctx context.Context, param types.LeaveChat) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/LeaveChat"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get up-to-date information about the chat. Returns a ChatFullInfo object on success.
// 
// https://core.telegram.org/bots/api#getchat
func (bot *Bot) GetChat(ctx context.Context, param types.GetChat) (*types.ChatFullInfo, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetChat"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.ChatFullInfo]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get a list of administrators in a chat, which aren&#39;t bots. Returns an Array of ChatMember objects.
// 
// https://core.telegram.org/bots/api#getchatadministrators
func (bot *Bot) GetChatAdministrators(ctx context.Context, param types.GetChatAdministrators) ([]types.ChatMember, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetChatAdministrators"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[[]types.ChatMember]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the number of members in a chat. Returns Int on success.
// 
// https://core.telegram.org/bots/api#getchatmembercount
func (bot *Bot) GetChatMemberCount(ctx context.Context, param types.GetChatMemberCount) (int64, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}
	
	url := URL + bot.token + "/GetChatMemberCount"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return 0, err
	}
	
	var result tgResponse[int64]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
// 
// https://core.telegram.org/bots/api#getchatmember
func (bot *Bot) GetChatMember(ctx context.Context, param types.GetChatMember) (*types.ChatMember, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetChatMember"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.ChatMember]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatstickerset
func (bot *Bot) SetChatStickerSet(ctx context.Context, param types.SetChatStickerSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetChatStickerSet"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletechatstickerset
func (bot *Bot) DeleteChatStickerSet(ctx context.Context, param types.DeleteChatStickerSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteChatStickerSet"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of Sticker objects.
// 
// https://core.telegram.org/bots/api#getforumtopiciconstickers
func (bot *Bot) GetForumTopicIconStickers(ctx context.Context, param types.GetForumTopicIconStickers) ([]types.Sticker, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetForumTopicIconStickers"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[[]types.Sticker]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to create a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns information about the created topic as a ForumTopic object.
// 
// https://core.telegram.org/bots/api#createforumtopic
func (bot *Bot) CreateForumTopic(ctx context.Context, param types.CreateForumTopic) (*types.ForumTopic, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/CreateForumTopic"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.ForumTopic]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
// 
// https://core.telegram.org/bots/api#editforumtopic
func (bot *Bot) EditForumTopic(ctx context.Context, param types.EditForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/EditForumTopic"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
// 
// https://core.telegram.org/bots/api#closeforumtopic
func (bot *Bot) CloseForumTopic(ctx context.Context, param types.CloseForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/CloseForumTopic"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
// 
// https://core.telegram.org/bots/api#reopenforumtopic
func (bot *Bot) ReopenForumTopic(ctx context.Context, param types.ReopenForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/ReopenForumTopic"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a forum topic along with all its messages in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#deleteforumtopic
func (bot *Bot) DeleteForumTopic(ctx context.Context, param types.DeleteForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteForumTopic"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
// 
// https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (bot *Bot) UnpinAllForumTopicMessages(ctx context.Context, param types.UnpinAllForumTopicMessages) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/UnpinAllForumTopicMessages"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit the name of the &#39;General&#39; topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#editgeneralforumtopic
func (bot *Bot) EditGeneralForumTopic(ctx context.Context, param types.EditGeneralForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/EditGeneralForumTopic"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to close an open &#39;General&#39; topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#closegeneralforumtopic
func (bot *Bot) CloseGeneralForumTopic(ctx context.Context, param types.CloseGeneralForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/CloseGeneralForumTopic"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to reopen a closed &#39;General&#39; topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
// 
// https://core.telegram.org/bots/api#reopengeneralforumtopic
func (bot *Bot) ReopenGeneralForumTopic(ctx context.Context, param types.ReopenGeneralForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/ReopenGeneralForumTopic"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to hide the &#39;General&#39; topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
// 
// https://core.telegram.org/bots/api#hidegeneralforumtopic
func (bot *Bot) HideGeneralForumTopic(ctx context.Context, param types.HideGeneralForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/HideGeneralForumTopic"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to unhide the &#39;General&#39; topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (bot *Bot) UnhideGeneralForumTopic(ctx context.Context, param types.UnhideGeneralForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/UnhideGeneralForumTopic"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
// 
// https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func (bot *Bot) UnpinAllGeneralForumTopicMessages(ctx context.Context, param types.UnpinAllGeneralForumTopicMessages) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/UnpinAllGeneralForumTopicMessages"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
//
// https://core.telegram.org/bots/api#answercallbackquery
func (bot *Bot) AnswerCallbackQuery(ctx context.Context, param types.AnswerCallbackQuery) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/AnswerCallbackQuery"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
// 
// https://core.telegram.org/bots/api#getuserchatboosts
func (bot *Bot) GetUserChatBoosts(ctx context.Context, param types.GetUserChatBoosts) (*types.UserChatBoosts, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetUserChatBoosts"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.UserChatBoosts]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get information about the connection of the bot with a business account. Returns a BusinessConnection object on success.
// 
// https://core.telegram.org/bots/api#getbusinessconnection
func (bot *Bot) GetBusinessConnection(ctx context.Context, param types.GetBusinessConnection) (*types.BusinessConnection, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetBusinessConnection"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.BusinessConnection]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the list of the bot&#39;s commands. See this manual for more details about bot commands. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmycommands
func (bot *Bot) SetMyCommands(ctx context.Context, param types.SetMyCommands) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetMyCommands"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete the list of the bot&#39;s commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletemycommands
func (bot *Bot) DeleteMyCommands(ctx context.Context, param types.DeleteMyCommands) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteMyCommands"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current list of the bot&#39;s commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren&#39;t set, an empty list is returned.
// 
// https://core.telegram.org/bots/api#getmycommands
func (bot *Bot) GetMyCommands(ctx context.Context, param types.GetMyCommands) ([]types.BotCommand, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetMyCommands"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[[]types.BotCommand]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the bot&#39;s name. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmyname
func (bot *Bot) SetMyName(ctx context.Context, param types.SetMyName) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetMyName"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current bot name for the given user language. Returns BotName on success.
// 
// https://core.telegram.org/bots/api#getmyname
func (bot *Bot) GetMyName(ctx context.Context, param types.GetMyName) (*types.BotName, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetMyName"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.BotName]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the bot&#39;s description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmydescription
func (bot *Bot) SetMyDescription(ctx context.Context, param types.SetMyDescription) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetMyDescription"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current bot description for the given user language. Returns BotDescription on success.
// 
// https://core.telegram.org/bots/api#getmydescription
func (bot *Bot) GetMyDescription(ctx context.Context, param types.GetMyDescription) (*types.BotDescription, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetMyDescription"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.BotDescription]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the bot&#39;s short description, which is shown on the bot&#39;s profile page and is sent together with the link when users share the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmyshortdescription
func (bot *Bot) SetMyShortDescription(ctx context.Context, param types.SetMyShortDescription) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetMyShortDescription"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current bot short description for the given user language. Returns BotShortDescription on success.
// 
// https://core.telegram.org/bots/api#getmyshortdescription
func (bot *Bot) GetMyShortDescription(ctx context.Context, param types.GetMyShortDescription) (*types.BotShortDescription, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetMyShortDescription"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.BotShortDescription]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the bot&#39;s menu button in a private chat, or the default menu button. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatmenubutton
func (bot *Bot) SetChatMenuButton(ctx context.Context, param types.SetChatMenuButton) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetChatMenuButton"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current value of the bot&#39;s menu button in a private chat, or the default menu button. Returns MenuButton on success.
// 
// https://core.telegram.org/bots/api#getchatmenubutton
func (bot *Bot) GetChatMenuButton(ctx context.Context, param types.GetChatMenuButton) (*types.MenuButton, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetChatMenuButton"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.MenuButton]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the default administrator rights requested by the bot when it&#39;s added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (bot *Bot) SetMyDefaultAdministratorRights(ctx context.Context, param types.SetMyDefaultAdministratorRights) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetMyDefaultAdministratorRights"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
// 
// https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (bot *Bot) GetMyDefaultAdministratorRights(ctx context.Context, param types.GetMyDefaultAdministratorRights) (*types.ChatAdministratorRights, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetMyDefaultAdministratorRights"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.ChatAdministratorRights]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Returns the list of gifts that can be sent by the bot to users and channel chats. Requires no parameters. Returns a Gifts object.
// 
// https://core.telegram.org/bots/api#getavailablegifts
func (bot *Bot) GetAvailableGifts(ctx context.Context, param types.GetAvailableGifts) (*types.Gifts, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetAvailableGifts"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Gifts]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Sends a gift to the given user or channel chat. The gift can&#39;t be converted to Telegram Stars by the receiver. Returns True on success.
// 
// https://core.telegram.org/bots/api#sendgift
func (bot *Bot) SendGift(ctx context.Context, param types.SendGift) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SendGift"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Gifts a Telegram Premium subscription to the given user. Returns True on success.
// 
// https://core.telegram.org/bots/api#giftpremiumsubscription
func (bot *Bot) GiftPremiumSubscription(ctx context.Context, param types.GiftPremiumSubscription) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/GiftPremiumSubscription"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Verifies a user on behalf of the organization which is represented by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#verifyuser
func (bot *Bot) VerifyUser(ctx context.Context, param types.VerifyUser) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/VerifyUser"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Verifies a chat on behalf of the organization which is represented by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#verifychat
func (bot *Bot) VerifyChat(ctx context.Context, param types.VerifyChat) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/VerifyChat"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Removes verification from a user who is currently verified on behalf of the organization represented by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#removeuserverification
func (bot *Bot) RemoveUserVerification(ctx context.Context, param types.RemoveUserVerification) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/RemoveUserVerification"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Removes verification from a chat that is currently verified on behalf of the organization represented by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#removechatverification
func (bot *Bot) RemoveChatVerification(ctx context.Context, param types.RemoveChatVerification) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/RemoveChatVerification"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Marks incoming message as read on behalf of a business account. Requires the can_read_messages business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#readbusinessmessage
func (bot *Bot) ReadBusinessMessage(ctx context.Context, param types.ReadBusinessMessage) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/ReadBusinessMessage"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Delete messages on behalf of a business account. Requires the can_delete_sent_messages business bot right to delete messages sent by the bot itself, or the can_delete_all_messages business bot right to delete any message. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletebusinessmessages
func (bot *Bot) DeleteBusinessMessages(ctx context.Context, param types.DeleteBusinessMessages) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteBusinessMessages"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the first and last name of a managed business account. Requires the can_change_name business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#setbusinessaccountname
func (bot *Bot) SetBusinessAccountName(ctx context.Context, param types.SetBusinessAccountName) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetBusinessAccountName"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the username of a managed business account. Requires the can_change_username business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#setbusinessaccountusername
func (bot *Bot) SetBusinessAccountUsername(ctx context.Context, param types.SetBusinessAccountUsername) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetBusinessAccountUsername"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the bio of a managed business account. Requires the can_change_bio business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#setbusinessaccountbio
func (bot *Bot) SetBusinessAccountBio(ctx context.Context, param types.SetBusinessAccountBio) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetBusinessAccountBio"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#setbusinessaccountprofilephoto
func (bot *Bot) SetBusinessAccountProfilePhoto(ctx context.Context, param types.SetBusinessAccountProfilePhoto) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetBusinessAccountProfilePhoto"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Removes the current profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#removebusinessaccountprofilephoto
func (bot *Bot) RemoveBusinessAccountProfilePhoto(ctx context.Context, param types.RemoveBusinessAccountProfilePhoto) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/RemoveBusinessAccountProfilePhoto"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the privacy settings pertaining to incoming gifts in a managed business account. Requires the can_change_gift_settings business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#setbusinessaccountgiftsettings
func (bot *Bot) SetBusinessAccountGiftSettings(ctx context.Context, param types.SetBusinessAccountGiftSettings) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetBusinessAccountGiftSettings"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Returns the amount of Telegram Stars owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns StarAmount on success.
// 
// https://core.telegram.org/bots/api#getbusinessaccountstarbalance
func (bot *Bot) GetBusinessAccountStarBalance(ctx context.Context, param types.GetBusinessAccountStarBalance) (*types.StarAmount, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetBusinessAccountStarBalance"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.StarAmount]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Transfers Telegram Stars from the business account balance to the bot&#39;s balance. Requires the can_transfer_stars business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#transferbusinessaccountstars
func (bot *Bot) TransferBusinessAccountStars(ctx context.Context, param types.TransferBusinessAccountStars) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/TransferBusinessAccountStars"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Returns the gifts received and owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns OwnedGifts on success.
// 
// https://core.telegram.org/bots/api#getbusinessaccountgifts
func (bot *Bot) GetBusinessAccountGifts(ctx context.Context, param types.GetBusinessAccountGifts) (*types.OwnedGifts, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetBusinessAccountGifts"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.OwnedGifts]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Converts a given regular gift to Telegram Stars. Requires the can_convert_gifts_to_stars business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#convertgifttostars
func (bot *Bot) ConvertGiftToStars(ctx context.Context, param types.ConvertGiftToStars) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/ConvertGiftToStars"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Upgrades a given regular gift to a unique gift. Requires the can_transfer_and_upgrade_gifts business bot right. Additionally requires the can_transfer_stars business bot right if the upgrade is paid. Returns True on success.
// 
// https://core.telegram.org/bots/api#upgradegift
func (bot *Bot) UpgradeGift(ctx context.Context, param types.UpgradeGift) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/UpgradeGift"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Transfers an owned unique gift to another user. Requires the can_transfer_and_upgrade_gifts business bot right. Requires can_transfer_stars business bot right if the transfer is paid. Returns True on success.
// 
// https://core.telegram.org/bots/api#transfergift
func (bot *Bot) TransferGift(ctx context.Context, param types.TransferGift) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/TransferGift"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Posts a story on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns Story on success.
// 
// https://core.telegram.org/bots/api#poststory
func (bot *Bot) PostStory(ctx context.Context, param types.PostStory) (*types.Story, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/PostStory"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Story]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Edits a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns Story on success.
// 
// https://core.telegram.org/bots/api#editstory
func (bot *Bot) EditStory(ctx context.Context, param types.EditStory) (*types.Story, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/EditStory"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Story]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Deletes a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletestory
func (bot *Bot) DeleteStory(ctx context.Context, param types.DeleteStory) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteStory"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
// 
// https://core.telegram.org/bots/api#editmessagetext
func (bot *Bot) EditMessageText(ctx context.Context, param types.EditMessageText) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/EditMessageText"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
// 
// https://core.telegram.org/bots/api#editmessagecaption
func (bot *Bot) EditMessageCaption(ctx context.Context, param types.EditMessageCaption) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/EditMessageCaption"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit animation, audio, document, photo, or video messages, or to add media to text messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can&#39;t be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
// 
// https://core.telegram.org/bots/api#editmessagemedia
func (bot *Bot) EditMessageMedia(ctx context.Context, param types.EditMessageMedia) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/EditMessageMedia"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
// 
// https://core.telegram.org/bots/api#editmessagelivelocation
func (bot *Bot) EditMessageLiveLocation(ctx context.Context, param types.EditMessageLiveLocation) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/EditMessageLiveLocation"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
// 
// https://core.telegram.org/bots/api#stopmessagelivelocation
func (bot *Bot) StopMessageLiveLocation(ctx context.Context, param types.StopMessageLiveLocation) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/StopMessageLiveLocation"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit a checklist on behalf of a connected business account. On success, the edited Message is returned.
// 
// https://core.telegram.org/bots/api#editmessagechecklist
func (bot *Bot) EditMessageChecklist(ctx context.Context, param types.EditMessageChecklist) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/EditMessageChecklist"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
// 
// https://core.telegram.org/bots/api#editmessagereplymarkup
func (bot *Bot) EditMessageReplyMarkup(ctx context.Context, param types.EditMessageReplyMarkup) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/EditMessageReplyMarkup"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
// 
// https://core.telegram.org/bots/api#stoppoll
func (bot *Bot) StopPoll(ctx context.Context, param types.StopPoll) (*types.Poll, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/StopPoll"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Poll]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to approve a suggested post in a direct messages chat. The bot must have the &#39;can_post_messages&#39; administrator right in the corresponding channel chat. Returns True on success.
// 
// https://core.telegram.org/bots/api#approvesuggestedpost
func (bot *Bot) ApproveSuggestedPost(ctx context.Context, param types.ApproveSuggestedPost) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/ApproveSuggestedPost"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to decline a suggested post in a direct messages chat. The bot must have the &#39;can_manage_direct_messages&#39; administrator right in the corresponding channel chat. Returns True on success.
// 
// https://core.telegram.org/bots/api#declinesuggestedpost
func (bot *Bot) DeclineSuggestedPost(ctx context.Context, param types.DeclineSuggestedPost) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeclineSuggestedPost"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Service messages about a supergroup, channel, or forum topic creation can&#39;t be deleted.- A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.- Bots can delete outgoing messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages administrator right in a supergroup or a channel, it can delete any message there.- If the bot has can_manage_direct_messages administrator right in a channel, it can delete any message in the corresponding direct messages chat.Returns True on success.
// 
// https://core.telegram.org/bots/api#deletemessage
func (bot *Bot) DeleteMessage(ctx context.Context, param types.DeleteMessage) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteMessage"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete multiple messages simultaneously. If some of the specified messages can&#39;t be found, they are skipped. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletemessages
func (bot *Bot) DeleteMessages(ctx context.Context, param types.DeleteMessages) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteMessages"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendsticker
func (bot *Bot) SendSticker(ctx context.Context, param types.SendSticker) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendSticker"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get a sticker set. On success, a StickerSet object is returned.
// 
// https://core.telegram.org/bots/api#getstickerset
func (bot *Bot) GetStickerSet(ctx context.Context, param types.GetStickerSet) (*types.StickerSet, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetStickerSet"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.StickerSet]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
// 
// https://core.telegram.org/bots/api#getcustomemojistickers
func (bot *Bot) GetCustomEmojiStickers(ctx context.Context, param types.GetCustomEmojiStickers) ([]types.Sticker, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetCustomEmojiStickers"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[[]types.Sticker]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to upload a file with a sticker for later use in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods (the file can be used multiple times). Returns the uploaded File on success.
// 
// https://core.telegram.org/bots/api#uploadstickerfile
func (bot *Bot) UploadStickerFile(ctx context.Context, param types.UploadStickerFile) (*types.File, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/UploadStickerFile"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.File]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
// 
// https://core.telegram.org/bots/api#createnewstickerset
func (bot *Bot) CreateNewStickerSet(ctx context.Context, param types.CreateNewStickerSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/CreateNewStickerSet"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.
// 
// https://core.telegram.org/bots/api#addstickertoset
func (bot *Bot) AddStickerToSet(ctx context.Context, param types.AddStickerToSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/AddStickerToSet"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickerpositioninset
func (bot *Bot) SetStickerPositionInSet(ctx context.Context, param types.SetStickerPositionInSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetStickerPositionInSet"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a sticker from a set created by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletestickerfromset
func (bot *Bot) DeleteStickerFromSet(ctx context.Context, param types.DeleteStickerFromSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteStickerFromSet"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet. Returns True on success.
// 
// https://core.telegram.org/bots/api#replacestickerinset
func (bot *Bot) ReplaceStickerInSet(ctx context.Context, param types.ReplaceStickerInSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/ReplaceStickerInSet"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickeremojilist
func (bot *Bot) SetStickerEmojiList(ctx context.Context, param types.SetStickerEmojiList) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetStickerEmojiList"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickerkeywords
func (bot *Bot) SetStickerKeywords(ctx context.Context, param types.SetStickerKeywords) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetStickerKeywords"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickermaskposition
func (bot *Bot) SetStickerMaskPosition(ctx context.Context, param types.SetStickerMaskPosition) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetStickerMaskPosition"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set the title of a created sticker set. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickersettitle
func (bot *Bot) SetStickerSetTitle(ctx context.Context, param types.SetStickerSetTitle) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetStickerSetTitle"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickersetthumbnail
func (bot *Bot) SetStickerSetThumbnail(ctx context.Context, param types.SetStickerSetThumbnail) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetStickerSetThumbnail"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.
// 
// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func (bot *Bot) SetCustomEmojiStickerSetThumbnail(ctx context.Context, param types.SetCustomEmojiStickerSetThumbnail) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetCustomEmojiStickerSetThumbnail"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a sticker set that was created by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletestickerset
func (bot *Bot) DeleteStickerSet(ctx context.Context, param types.DeleteStickerSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/DeleteStickerSet"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
// 
// https://core.telegram.org/bots/api#answerinlinequery
func (bot *Bot) AnswerInlineQuery(ctx context.Context, param types.AnswerInlineQuery) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/AnswerInlineQuery"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
// 
// https://core.telegram.org/bots/api#answerwebappquery
func (bot *Bot) AnswerWebAppQuery(ctx context.Context, param types.AnswerWebAppQuery) (*types.SentWebAppMessage, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/AnswerWebAppQuery"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.SentWebAppMessage]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Stores a message that can be sent by a user of a Mini App. Returns a PreparedInlineMessage object.
// 
// https://core.telegram.org/bots/api#savepreparedinlinemessage
func (bot *Bot) SavePreparedInlineMessage(ctx context.Context, param types.SavePreparedInlineMessage) (*types.PreparedInlineMessage, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SavePreparedInlineMessage"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.PreparedInlineMessage]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send invoices. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendinvoice
func (bot *Bot) SendInvoice(ctx context.Context, param types.SendInvoice) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendInvoice"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to create a link for an invoice. Returns the created invoice link as String on success.
// 
// https://core.telegram.org/bots/api#createinvoicelink
func (bot *Bot) CreateInvoiceLink(ctx context.Context, param types.CreateInvoiceLink) (string, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	
	url := URL + bot.token + "/CreateInvoiceLink"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return "", err
	}
	
	var result tgResponse[string]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
// 
// https://core.telegram.org/bots/api#answershippingquery
func (bot *Bot) AnswerShippingQuery(ctx context.Context, param types.AnswerShippingQuery) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/AnswerShippingQuery"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
// 
// https://core.telegram.org/bots/api#answerprecheckoutquery
func (bot *Bot) AnswerPreCheckoutQuery(ctx context.Context, param types.AnswerPreCheckoutQuery) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/AnswerPreCheckoutQuery"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// A method to get the current Telegram Stars balance of the bot. Requires no parameters. On success, returns a StarAmount object.
// 
// https://core.telegram.org/bots/api#getmystarbalance
func (bot *Bot) GetMyStarBalance(ctx context.Context, param types.GetMyStarBalance) (*types.StarAmount, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetMyStarBalance"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.StarAmount]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Returns the bot&#39;s Telegram Star transactions in chronological order. On success, returns a StarTransactions object.
// 
// https://core.telegram.org/bots/api#getstartransactions
func (bot *Bot) GetStarTransactions(ctx context.Context, param types.GetStarTransactions) (*types.StarTransactions, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetStarTransactions"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.StarTransactions]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Refunds a successful payment in Telegram Stars. Returns True on success.
// 
// https://core.telegram.org/bots/api#refundstarpayment
func (bot *Bot) RefundStarPayment(ctx context.Context, param types.RefundStarPayment) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/RefundStarPayment"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Returns True on success.
// 
// https://core.telegram.org/bots/api#edituserstarsubscription
func (bot *Bot) EditUserStarSubscription(ctx context.Context, param types.EditUserStarSubscription) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/EditUserStarSubscription"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
// 
// https://core.telegram.org/bots/api#setpassportdataerrors
func (bot *Bot) SetPassportDataErrors(ctx context.Context, param types.SetPassportDataErrors) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + bot.token + "/SetPassportDataErrors"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return false, err
	}
	
	var result tgResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send a game. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendgame
func (bot *Bot) SendGame(ctx context.Context, param types.SendGame) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SendGame"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user&#39;s current score in the chat and force is False.
// 
// https://core.telegram.org/bots/api#setgamescore
func (bot *Bot) SetGameScore(ctx context.Context, param types.SetGameScore) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/SetGameScore"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
//
// This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.
//
// https://core.telegram.org/bots/api#getgamehighscores
func (bot *Bot) GetGameHighScores(ctx context.Context, param types.GetGameHighScores) ([]types.GameHighScore, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + bot.token + "/GetGameHighScores"
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	
	var result tgResponse[[]types.GameHighScore]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

